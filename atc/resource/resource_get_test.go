package resource_test

import (
	"context"
	"errors"

	"github.com/concourse/concourse/atc"
	"github.com/concourse/concourse/atc/resource"
	"github.com/concourse/concourse/atc/runtime"
	"github.com/concourse/concourse/atc/runtime/runtimefakes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
)

var _ = Describe("Resource Get", func() {
	var (
		ctx             context.Context
		someProcessSpec runtime.ProcessSpec
		fakeRunnable    runtimefakes.FakeRunner

		getVersionResult runtime.VersionResult

		source  atc.Source
		params  atc.Params
		version atc.Version

		resource resource.Resource

		getErr error
	)

	BeforeEach(func() {
		ctx = context.Background()

		source = atc.Source{"some": "source"}
		version = atc.Version{"some": "version"}
		params = atc.Params{"some": "params"}

		someProcessSpec.Path = "some/fake/path"
		someProcessSpec.Args = []string{"first-arg", "some-other-arg"}
		someProcessSpec.StderrWriter = gbytes.NewBuffer()

		resource = resourceFactory.NewResource(source, params, version)

	})

	JustBeforeEach(func() {
		getVersionResult, getErr = resource.Get(ctx, someProcessSpec, &fakeRunnable)
	})

	Context("when Runnable -> RunScript succeeds", func() {
		BeforeEach(func() {
			fakeRunnable.RunScriptReturns(nil)
		})

		It("Invokes Runnable -> RunScript with the correct arguments", func() {
			actualCtx, actualSpecPath, actualArgs,
				actualInput, actualVersionResultRef, actualSpecStdErrWriter,
				actualRecoverableBool := fakeRunnable.RunScriptArgsForCall(0)

			signature, err := resource.Signature()
			Expect(err).ToNot(HaveOccurred())

			Expect(actualCtx).To(Equal(ctx))
			Expect(actualSpecPath).To(Equal(someProcessSpec.Path))
			Expect(actualArgs).To(Equal(someProcessSpec.Args))
			Expect(actualInput).To(Equal(signature))
			Expect(actualVersionResultRef).To(Equal(&getVersionResult))
			Expect(actualSpecStdErrWriter).To(Equal(someProcessSpec.StderrWriter))
			Expect(actualRecoverableBool).To(BeTrue())
		})

		It("doesnt return an error", func() {
			Expect(getErr).To(BeNil())
		})
	})

	Context("when Runnable -> RunScript returns an error", func() {
		var disasterErr = errors.New("there was an issue")
		BeforeEach(func() {
			fakeRunnable.RunScriptReturns(disasterErr)
		})
		It("returns the error", func() {
			Expect(getErr).To(Equal(disasterErr))
		})
	})

})
