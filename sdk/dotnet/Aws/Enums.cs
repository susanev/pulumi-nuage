// *** WARNING: this file was generated by Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.Nuage.Aws
{
    /// <summary>
    /// Architecture, either 'X86_64' or 'arm64'.
    /// </summary>
    [EnumType]
    public readonly struct ArchitectureType : IEquatable<ArchitectureType>
    {
        private readonly string _value;

        private ArchitectureType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// X86_64 architecture.
        /// </summary>
        public static ArchitectureType X86_64 { get; } = new ArchitectureType("X86_64");
        /// <summary>
        /// ARM64 architecture.
        /// </summary>
        public static ArchitectureType ARM64 { get; } = new ArchitectureType("ARM64");

        public static bool operator ==(ArchitectureType left, ArchitectureType right) => left.Equals(right);
        public static bool operator !=(ArchitectureType left, ArchitectureType right) => !left.Equals(right);

        public static explicit operator string(ArchitectureType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ArchitectureType other && Equals(other);
        public bool Equals(ArchitectureType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}