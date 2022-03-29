package plugins_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

  "os"

	. "github.com/bytesparadise/libasciidoc/testsupport"
	"github.com/bytesparadise/libasciidoc/pkg/plugins"
	"github.com/bytesparadise/libasciidoc/pkg/configuration"
)

var _ = Describe("Plugins", func() {
  Describe("Using plugins", func() {
    Context("that pass through", func() {
      It("should render demo.adoc the same as without the plugin", func() {
				// given
				_, err := os.Stat("../../test/compat/demo.adoc")
				Expect(err).NotTo(HaveOccurred())
				_, err = os.Stat("../../test/plugins/pass")
				Expect(err).NotTo(HaveOccurred())
        plugins, err := plugins.LoadPlugins([]string{"../../test/plugins/pass"})
        Expect(err).NotTo(HaveOccurred())

				// when
				original, _, err := RenderHTMLFromFile("../../test/compat/demo.adoc")
        Expect(err).NotTo(HaveOccurred())
				withPlugin, _, err := RenderHTMLFromFile("../../test/compat/demo.adoc",
			    configuration.WithPlugins(plugins))
        Expect(err).NotTo(HaveOccurred())

        // then
        Expect(original).To(Equal(withPlugin))
      })
    })
    /*Context("that cause errors", func() {
    })
    Context("that create new documents", func() {
    })*/
  })
})
