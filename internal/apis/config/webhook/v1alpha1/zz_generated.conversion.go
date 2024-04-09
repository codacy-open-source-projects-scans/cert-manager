//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright The cert-manager Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by conversion-gen. DO NOT EDIT.

package v1alpha1

import (
	time "time"
	unsafe "unsafe"

	webhook "github.com/cert-manager/cert-manager/internal/apis/config/webhook"
	v1alpha1 "github.com/cert-manager/cert-manager/pkg/apis/config/webhook/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*v1alpha1.DynamicServingConfig)(nil), (*webhook.DynamicServingConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_DynamicServingConfig_To_webhook_DynamicServingConfig(a.(*v1alpha1.DynamicServingConfig), b.(*webhook.DynamicServingConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*webhook.DynamicServingConfig)(nil), (*v1alpha1.DynamicServingConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_webhook_DynamicServingConfig_To_v1alpha1_DynamicServingConfig(a.(*webhook.DynamicServingConfig), b.(*v1alpha1.DynamicServingConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha1.FilesystemServingConfig)(nil), (*webhook.FilesystemServingConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_FilesystemServingConfig_To_webhook_FilesystemServingConfig(a.(*v1alpha1.FilesystemServingConfig), b.(*webhook.FilesystemServingConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*webhook.FilesystemServingConfig)(nil), (*v1alpha1.FilesystemServingConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_webhook_FilesystemServingConfig_To_v1alpha1_FilesystemServingConfig(a.(*webhook.FilesystemServingConfig), b.(*v1alpha1.FilesystemServingConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha1.TLSConfig)(nil), (*webhook.TLSConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_TLSConfig_To_webhook_TLSConfig(a.(*v1alpha1.TLSConfig), b.(*webhook.TLSConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*webhook.TLSConfig)(nil), (*v1alpha1.TLSConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_webhook_TLSConfig_To_v1alpha1_TLSConfig(a.(*webhook.TLSConfig), b.(*v1alpha1.TLSConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha1.WebhookConfiguration)(nil), (*webhook.WebhookConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_WebhookConfiguration_To_webhook_WebhookConfiguration(a.(*v1alpha1.WebhookConfiguration), b.(*webhook.WebhookConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*webhook.WebhookConfiguration)(nil), (*v1alpha1.WebhookConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_webhook_WebhookConfiguration_To_v1alpha1_WebhookConfiguration(a.(*webhook.WebhookConfiguration), b.(*v1alpha1.WebhookConfiguration), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1alpha1_DynamicServingConfig_To_webhook_DynamicServingConfig(in *v1alpha1.DynamicServingConfig, out *webhook.DynamicServingConfig, s conversion.Scope) error {
	out.SecretNamespace = in.SecretNamespace
	out.SecretName = in.SecretName
	out.DNSNames = *(*[]string)(unsafe.Pointer(&in.DNSNames))
	out.LeafDuration = time.Duration(in.LeafDuration)
	return nil
}

// Convert_v1alpha1_DynamicServingConfig_To_webhook_DynamicServingConfig is an autogenerated conversion function.
func Convert_v1alpha1_DynamicServingConfig_To_webhook_DynamicServingConfig(in *v1alpha1.DynamicServingConfig, out *webhook.DynamicServingConfig, s conversion.Scope) error {
	return autoConvert_v1alpha1_DynamicServingConfig_To_webhook_DynamicServingConfig(in, out, s)
}

func autoConvert_webhook_DynamicServingConfig_To_v1alpha1_DynamicServingConfig(in *webhook.DynamicServingConfig, out *v1alpha1.DynamicServingConfig, s conversion.Scope) error {
	out.SecretNamespace = in.SecretNamespace
	out.SecretName = in.SecretName
	out.DNSNames = *(*[]string)(unsafe.Pointer(&in.DNSNames))
	out.LeafDuration = time.Duration(in.LeafDuration)
	return nil
}

// Convert_webhook_DynamicServingConfig_To_v1alpha1_DynamicServingConfig is an autogenerated conversion function.
func Convert_webhook_DynamicServingConfig_To_v1alpha1_DynamicServingConfig(in *webhook.DynamicServingConfig, out *v1alpha1.DynamicServingConfig, s conversion.Scope) error {
	return autoConvert_webhook_DynamicServingConfig_To_v1alpha1_DynamicServingConfig(in, out, s)
}

func autoConvert_v1alpha1_FilesystemServingConfig_To_webhook_FilesystemServingConfig(in *v1alpha1.FilesystemServingConfig, out *webhook.FilesystemServingConfig, s conversion.Scope) error {
	out.CertFile = in.CertFile
	out.KeyFile = in.KeyFile
	return nil
}

// Convert_v1alpha1_FilesystemServingConfig_To_webhook_FilesystemServingConfig is an autogenerated conversion function.
func Convert_v1alpha1_FilesystemServingConfig_To_webhook_FilesystemServingConfig(in *v1alpha1.FilesystemServingConfig, out *webhook.FilesystemServingConfig, s conversion.Scope) error {
	return autoConvert_v1alpha1_FilesystemServingConfig_To_webhook_FilesystemServingConfig(in, out, s)
}

func autoConvert_webhook_FilesystemServingConfig_To_v1alpha1_FilesystemServingConfig(in *webhook.FilesystemServingConfig, out *v1alpha1.FilesystemServingConfig, s conversion.Scope) error {
	out.CertFile = in.CertFile
	out.KeyFile = in.KeyFile
	return nil
}

// Convert_webhook_FilesystemServingConfig_To_v1alpha1_FilesystemServingConfig is an autogenerated conversion function.
func Convert_webhook_FilesystemServingConfig_To_v1alpha1_FilesystemServingConfig(in *webhook.FilesystemServingConfig, out *v1alpha1.FilesystemServingConfig, s conversion.Scope) error {
	return autoConvert_webhook_FilesystemServingConfig_To_v1alpha1_FilesystemServingConfig(in, out, s)
}

func autoConvert_v1alpha1_TLSConfig_To_webhook_TLSConfig(in *v1alpha1.TLSConfig, out *webhook.TLSConfig, s conversion.Scope) error {
	out.CipherSuites = *(*[]string)(unsafe.Pointer(&in.CipherSuites))
	out.MinTLSVersion = in.MinTLSVersion
	if err := Convert_v1alpha1_FilesystemServingConfig_To_webhook_FilesystemServingConfig(&in.Filesystem, &out.Filesystem, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_DynamicServingConfig_To_webhook_DynamicServingConfig(&in.Dynamic, &out.Dynamic, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_TLSConfig_To_webhook_TLSConfig is an autogenerated conversion function.
func Convert_v1alpha1_TLSConfig_To_webhook_TLSConfig(in *v1alpha1.TLSConfig, out *webhook.TLSConfig, s conversion.Scope) error {
	return autoConvert_v1alpha1_TLSConfig_To_webhook_TLSConfig(in, out, s)
}

func autoConvert_webhook_TLSConfig_To_v1alpha1_TLSConfig(in *webhook.TLSConfig, out *v1alpha1.TLSConfig, s conversion.Scope) error {
	out.CipherSuites = *(*[]string)(unsafe.Pointer(&in.CipherSuites))
	out.MinTLSVersion = in.MinTLSVersion
	if err := Convert_webhook_FilesystemServingConfig_To_v1alpha1_FilesystemServingConfig(&in.Filesystem, &out.Filesystem, s); err != nil {
		return err
	}
	if err := Convert_webhook_DynamicServingConfig_To_v1alpha1_DynamicServingConfig(&in.Dynamic, &out.Dynamic, s); err != nil {
		return err
	}
	return nil
}

// Convert_webhook_TLSConfig_To_v1alpha1_TLSConfig is an autogenerated conversion function.
func Convert_webhook_TLSConfig_To_v1alpha1_TLSConfig(in *webhook.TLSConfig, out *v1alpha1.TLSConfig, s conversion.Scope) error {
	return autoConvert_webhook_TLSConfig_To_v1alpha1_TLSConfig(in, out, s)
}

func autoConvert_v1alpha1_WebhookConfiguration_To_webhook_WebhookConfiguration(in *v1alpha1.WebhookConfiguration, out *webhook.WebhookConfiguration, s conversion.Scope) error {
	if err := v1.Convert_Pointer_int32_To_int32(&in.SecurePort, &out.SecurePort, s); err != nil {
		return err
	}
	if err := v1.Convert_Pointer_int32_To_int32(&in.HealthzPort, &out.HealthzPort, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_TLSConfig_To_webhook_TLSConfig(&in.TLSConfig, &out.TLSConfig, s); err != nil {
		return err
	}
	out.KubeConfig = in.KubeConfig
	out.APIServerHost = in.APIServerHost
	out.EnablePprof = in.EnablePprof
	out.PprofAddress = in.PprofAddress
	out.Logging = in.Logging
	out.FeatureGates = *(*map[string]bool)(unsafe.Pointer(&in.FeatureGates))
	return nil
}

// Convert_v1alpha1_WebhookConfiguration_To_webhook_WebhookConfiguration is an autogenerated conversion function.
func Convert_v1alpha1_WebhookConfiguration_To_webhook_WebhookConfiguration(in *v1alpha1.WebhookConfiguration, out *webhook.WebhookConfiguration, s conversion.Scope) error {
	return autoConvert_v1alpha1_WebhookConfiguration_To_webhook_WebhookConfiguration(in, out, s)
}

func autoConvert_webhook_WebhookConfiguration_To_v1alpha1_WebhookConfiguration(in *webhook.WebhookConfiguration, out *v1alpha1.WebhookConfiguration, s conversion.Scope) error {
	if err := v1.Convert_int32_To_Pointer_int32(&in.SecurePort, &out.SecurePort, s); err != nil {
		return err
	}
	if err := v1.Convert_int32_To_Pointer_int32(&in.HealthzPort, &out.HealthzPort, s); err != nil {
		return err
	}
	if err := Convert_webhook_TLSConfig_To_v1alpha1_TLSConfig(&in.TLSConfig, &out.TLSConfig, s); err != nil {
		return err
	}
	out.KubeConfig = in.KubeConfig
	out.APIServerHost = in.APIServerHost
	out.EnablePprof = in.EnablePprof
	out.PprofAddress = in.PprofAddress
	out.Logging = in.Logging
	out.FeatureGates = *(*map[string]bool)(unsafe.Pointer(&in.FeatureGates))
	return nil
}

// Convert_webhook_WebhookConfiguration_To_v1alpha1_WebhookConfiguration is an autogenerated conversion function.
func Convert_webhook_WebhookConfiguration_To_v1alpha1_WebhookConfiguration(in *webhook.WebhookConfiguration, out *v1alpha1.WebhookConfiguration, s conversion.Scope) error {
	return autoConvert_webhook_WebhookConfiguration_To_v1alpha1_WebhookConfiguration(in, out, s)
}
