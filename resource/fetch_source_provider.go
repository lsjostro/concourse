package resource

import (
	"os"
	"time"

	"code.cloudfoundry.org/lager"
	"github.com/concourse/atc"
	"github.com/concourse/atc/worker"
)

//go:generate counterfeiter . FetchSourceProviderFactory

type FetchSourceProviderFactory interface {
	NewFetchSourceProvider(
		logger lager.Logger,
		session Session,
		tags atc.Tags,
		teamID int,
		resourceTypes atc.ResourceTypes,
		resourceInstance ResourceInstance,
		resourceOptions ResourceOptions,
		containerCreator FetchContainerCreator,
	) FetchSourceProvider
}

//go:generate counterfeiter . FetchSourceProvider

type FetchSourceProvider interface {
	Get() (FetchSource, error)
}

//go:generate counterfeiter . FetchSource

type FetchSource interface {
	IsInitialized() (bool, error)
	LockName() (string, error)
	VersionedSource() VersionedSource
	Initialize(signals <-chan os.Signal, ready chan<- struct{}) error
	Release(*time.Duration)
}

type fetchSourceProviderFactory struct {
	workerClient worker.Client
}

func NewFetchSourceProviderFactory(workerClient worker.Client) FetchSourceProviderFactory {
	return &fetchSourceProviderFactory{
		workerClient: workerClient,
	}
}

func (f *fetchSourceProviderFactory) NewFetchSourceProvider(
	logger lager.Logger,
	session Session,
	tags atc.Tags,
	teamID int,
	resourceTypes atc.ResourceTypes,
	resourceInstance ResourceInstance,
	resourceOptions ResourceOptions,
	containerCreator FetchContainerCreator,
) FetchSourceProvider {
	return &fetchSourceProvider{
		logger:           logger,
		session:          session,
		tags:             tags,
		teamID:           teamID,
		resourceTypes:    resourceTypes,
		resourceInstance: resourceInstance,
		resourceOptions:  resourceOptions,
		containerCreator: containerCreator,
		workerClient:     f.workerClient,
	}
}

type fetchSourceProvider struct {
	logger           lager.Logger
	session          Session
	tags             atc.Tags
	teamID           int
	resourceTypes    atc.ResourceTypes
	resourceInstance ResourceInstance
	resourceOptions  ResourceOptions
	workerClient     worker.Client
	containerCreator FetchContainerCreator
}

func (f *fetchSourceProvider) Get() (FetchSource, error) {
	container, found, err := f.workerClient.FindContainerForIdentifier(f.logger, f.session.ID)
	if err != nil {
		f.logger.Error("failed-to-look-for-existing-container", err)
		return nil, err
	}

	if found {
		cacheVolume, cacheVolumeFound := findCacheVolumeForContainer(container)
		if cacheVolumeFound {
			return NewContainerFetchSource(f.logger, container, cacheVolume, f.resourceOptions), nil
		}
	}

	resourceSpec := worker.WorkerSpec{
		ResourceType: string(f.resourceOptions.ResourceType()),
		Tags:         f.tags,
		TeamID:       f.teamID,
	}

	chosenWorker, err := f.workerClient.Satisfying(resourceSpec, f.resourceTypes)
	if err != nil {
		f.logger.Error("no-workers-satisfying-spec", err)
		return nil, err
	}

	cachedVolume, cacheFound, err := f.resourceInstance.FindOn(f.logger, chosenWorker)
	if err != nil {
		f.logger.Error("failed-to-look-for-cache", err)
		return nil, err
	}

	if cacheFound {
		f.logger.Debug("found-cache", lager.Data{"volume": cachedVolume.Handle()})
	} else {
		f.logger.Debug("no-cache-found")

		cachedVolume, err = f.resourceInstance.FindOrCreateOn(f.logger, chosenWorker)
		if err != nil {
			f.logger.Error("failed-to-create-cache", err)
			return nil, err
		}
	}

	return NewVolumeFetchSource(
		f.logger,
		cachedVolume,
		chosenWorker,
		f.resourceOptions,
		f.containerCreator,
	), nil
}

func findCacheVolumeForContainer(container worker.Container) (worker.Volume, bool) {
	for _, mount := range container.VolumeMounts() {
		if mount.MountPath == ResourcesDir("get") {
			return mount.Volume, true
		}
	}

	return nil, false
}
