package k8s

import (
	configv1 "github.com/ericchiang/k8s/config/v1"
)

// Config is provided for backward compatibility with <= v1.2.0 of of the library
type Config = configv1.Config

// Preferences is provided for backward compatibility with <= v1.2.0 of of the library
type Preferences = configv1.Preferences

// Cluster is provided for backward compatibility with <= v1.2.0 of of the library
type Cluster = configv1.Cluster

// AuthInfo is provided for backward compatibility with <= v1.2.0 of of the library
type AuthInfo = configv1.AuthInfo

// Context is provided for backward compatibility with <= v1.2.0 of of the library
type Context = configv1.Context

// NamedCluster is provided for backward compatibility with <= v1.2.0 of of the library
type NamedCluster = configv1.NamedCluster

// NamedContext is provided for backward compatibility with <= v1.2.0 of of the library
type NamedContext = configv1.NamedContext

// NamedAuthInfo is provided for backward compatibility with <= v1.2.0 of of the library
type NamedAuthInfo = configv1.NamedAuthInfo

// NamedExtension is provided for backward compatibility with <= v1.2.0 of of the library
type NamedExtension = configv1.NamedExtension

// AuthProviderConfig is provided for backward compatibility with <= v1.2.0 of of the library
type AuthProviderConfig = configv1.AuthProviderConfig
