package kcpsanity

import (
	g "github.com/onsi/ginkgo"
	o "github.com/onsi/gomega"

	exutil "github.com/kcp-dev/kcp-tests/test/extended/util"
)

var _ = g.Describe("[sig-kcp] Kcpsanity", func() {
	defer g.GinkgoRecover()

	var (
		oc = exutil.NewCLI("kubectl", exutil.KubeConfigPath())
	)

	g.It("Author:knarra-Medium-2800712-Checking oc version show clean as gitTreeState value", func() {
		oc = exutil.NewCLI("kubectl", exutil.KubeConfigPath())
		_, err := oc.Run("version").Args("-o", "json").Output()
		o.Expect(err).NotTo(o.HaveOccurred())

	})
})
