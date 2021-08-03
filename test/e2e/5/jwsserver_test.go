// +build !unit

package e2e

import (
	"testing"

	framework "github.com/operator-framework/operator-sdk/pkg/test"
	"github.com/web-servers/jws-operator/pkg/apis"
	webserversv1alpha1 "github.com/web-servers/jws-operator/pkg/apis/webservers/v1alpha1"
	webserversframework "github.com/web-servers/jws-operator/test/framework"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestWebServer54(t *testing.T) {
	webServerList := &webserversv1alpha1.WebServerList{
		TypeMeta: metav1.TypeMeta{
			Kind:       "WebServer",
			APIVersion: "web.servers.org/v1alpha1",
		},
	}
	err := framework.AddToFrameworkScheme(apis.AddToScheme, webServerList)
	if err != nil {
		t.Fatalf("failed to add custom resource scheme to framework: %v", err)
	}
	// run subtests
	t.Run("ApplicationImageBasicTest", webServerApplicationImageBasicTest)
	t.Run("ApplicationImageScaleTest", webServerApplicationImageScaleTest)
	t.Run("ApplicationImageUpdateTest", webServerApplicationImageUpdateTest)
	t.Run("ApplicationImageSourcesBasicTest", webServerApplicationImageSourcesBasicTest)
	t.Run("ApplicationImageSourcesScaleTest", webServerApplicationImageSourcesScaleTest)
	t.Run("ImageStreamBasicTest", webServerImageStreamBasicTest)
	t.Run("ImageStreamScaleTest", webServerImageStreamScaleTest)
	t.Run("ImageStreamSourcesBasicTest", webServerImageStreamSourcesBasicTest)
	t.Run("ImageStreamSourcesScaleTest", webServerImageStreamSourcesScaleTest)
}

func webServerApplicationImageBasicTest(t *testing.T) {
	webserversframework.WebServerApplicationImageBasicTest(t, "quay.io/jfclere/tomcat10:latest", "/health")
}

func webServerApplicationImageScaleTest(t *testing.T) {
	webserversframework.WebServerApplicationImageScaleTest(t, "quay.io/jfclere/tomcat10:latest", "/health")
}

func webServerApplicationImageUpdateTest(t *testing.T) {
	webserversframework.WebServerApplicationImageUpdateTest(t, "quay.io/jfclere/tomcat10:latest", "quay.io/pitprok/tomcat10:latest", "/health")
}

func webServerApplicationImageSourcesBasicTest(t *testing.T) {
	webserversframework.WebServerApplicationImageSourcesBasicTest(t, "quay.io/jfclere/tomcat10:latest", "https://github.com/jfclere/demo-webapp", "jakartaEE", "/demo")
}

func webServerApplicationImageSourcesScaleTest(t *testing.T) {
	webserversframework.WebServerApplicationImageSourcesScaleTest(t, "quay.io/jfclere/tomcat10:latest", "https://github.com/jfclere/demo-webapp", "jakartaEE", "/demo")
}

func webServerImageStreamBasicTest(t *testing.T) {
	webserversframework.WebServerImageStreamBasicTest(t, "jboss-webserver54-openjdk8-tomcat9-ubi8-openshift", "/health")
}

func webServerImageStreamScaleTest(t *testing.T) {
	webserversframework.WebServerImageStreamScaleTest(t, "jboss-webserver54-openjdk8-tomcat9-ubi8-openshift", "/health")
}

func webServerImageStreamSourcesBasicTest(t *testing.T) {
	webserversframework.WebServerImageStreamSourcesBasicTest(t, "jboss-webserver54-openjdk8-tomcat9-ubi8-openshift", "https://github.com/jfclere/demo-webapp", "", "/demo-1.0/demo")
}

func webServerImageStreamSourcesScaleTest(t *testing.T) {
	webserversframework.WebServerImageStreamSourcesScaleTest(t, "jboss-webserver54-openjdk8-tomcat9-ubi8-openshift", "https://github.com/jfclere/demo-webapp", "", "/demo-1.0/demo")
}
