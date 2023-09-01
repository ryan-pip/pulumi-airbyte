// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Immutable;

namespace Pulumi.Airbyte
{
    public static class Config
    {
        [System.Diagnostics.CodeAnalysis.SuppressMessage("Microsoft.Design", "IDE1006", Justification = 
        "Double underscore prefix used to avoid conflicts with variable names.")]
        private sealed class __Value<T>
        {
            private readonly Func<T> _getter;
            private T _value = default!;
            private bool _set;

            public __Value(Func<T> getter)
            {
                _getter = getter;
            }

            public T Get() => _set ? _value : _getter();

            public void Set(T value)
            {
                _value = value;
                _set = true;
            }
        }

        private static readonly global::Pulumi.Config __config = new global::Pulumi.Config("airbyte");

        private static readonly __Value<string?> _bearerAuth = new __Value<string?>(() => __config.Get("bearerAuth"));
        public static string? BearerAuth
        {
            get => _bearerAuth.Get();
            set => _bearerAuth.Set(value);
        }

        private static readonly __Value<string?> _password = new __Value<string?>(() => __config.Get("password"));
        public static string? Password
        {
            get => _password.Get();
            set => _password.Set(value);
        }

        private static readonly __Value<string?> _serverUrl = new __Value<string?>(() => __config.Get("serverUrl"));
        /// <summary>
        /// Server URL (defaults to https://api.airbyte.com/v1)
        /// </summary>
        public static string? ServerUrl
        {
            get => _serverUrl.Get();
            set => _serverUrl.Set(value);
        }

        private static readonly __Value<string?> _username = new __Value<string?>(() => __config.Get("username"));
        public static string? Username
        {
            get => _username.Get();
            set => _username.Set(value);
        }

    }
}