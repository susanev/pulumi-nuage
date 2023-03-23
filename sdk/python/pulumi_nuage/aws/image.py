# coding=utf-8
# *** WARNING: this file was generated by Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['ImageArgs', 'Image']

@pulumi.input_type
class ImageArgs:
    def __init__(__self__, *,
                 dockerfile: pulumi.Input[str],
                 repository_url: pulumi.Input[str],
                 architecture: Optional[pulumi.Input[str]] = None,
                 context: Optional[pulumi.Input[str]] = None,
                 target: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Image resource.
        :param pulumi.Input[str] dockerfile: The path to the Dockerfile to use.
        :param pulumi.Input[str] repository_url: Url of the repository.
        :param pulumi.Input[str] architecture: Architecture, either `X86_64` or `ARM64`. Defaults to `X86_64`
        :param pulumi.Input[str] context: The path to the build context to use.
        :param pulumi.Input[str] target: The target of the Dockerfile to build
        """
        pulumi.set(__self__, "dockerfile", dockerfile)
        pulumi.set(__self__, "repository_url", repository_url)
        if architecture is not None:
            pulumi.set(__self__, "architecture", architecture)
        if context is not None:
            pulumi.set(__self__, "context", context)
        if target is not None:
            pulumi.set(__self__, "target", target)

    @property
    @pulumi.getter
    def dockerfile(self) -> pulumi.Input[str]:
        """
        The path to the Dockerfile to use.
        """
        return pulumi.get(self, "dockerfile")

    @dockerfile.setter
    def dockerfile(self, value: pulumi.Input[str]):
        pulumi.set(self, "dockerfile", value)

    @property
    @pulumi.getter(name="repositoryUrl")
    def repository_url(self) -> pulumi.Input[str]:
        """
        Url of the repository.
        """
        return pulumi.get(self, "repository_url")

    @repository_url.setter
    def repository_url(self, value: pulumi.Input[str]):
        pulumi.set(self, "repository_url", value)

    @property
    @pulumi.getter
    def architecture(self) -> Optional[pulumi.Input[str]]:
        """
        Architecture, either `X86_64` or `ARM64`. Defaults to `X86_64`
        """
        return pulumi.get(self, "architecture")

    @architecture.setter
    def architecture(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "architecture", value)

    @property
    @pulumi.getter
    def context(self) -> Optional[pulumi.Input[str]]:
        """
        The path to the build context to use.
        """
        return pulumi.get(self, "context")

    @context.setter
    def context(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "context", value)

    @property
    @pulumi.getter
    def target(self) -> Optional[pulumi.Input[str]]:
        """
        The target of the Dockerfile to build
        """
        return pulumi.get(self, "target")

    @target.setter
    def target(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "target", value)


class Image(pulumi.ComponentResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 architecture: Optional[pulumi.Input[str]] = None,
                 context: Optional[pulumi.Input[str]] = None,
                 dockerfile: Optional[pulumi.Input[str]] = None,
                 repository_url: Optional[pulumi.Input[str]] = None,
                 target: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to manage build and deployment of Docker builds. It automatically builds the Docker image and pushes it to the specified repository.

        ## Example Usage
        ### Basic Example

        ```python
        import pulumi_nuage as nuage
        
        repository = nuage.aws.Repository(
            "foo",
            name="repository",
            expire_in_days=30,
        )
        
        image = nuage.aws.Image(
            "foo",
            dockerfile="../api/Dockerfile",
            context="../",
            repository_url=repository.url,
        )
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] architecture: Architecture, either `X86_64` or `ARM64`. Defaults to `X86_64`
        :param pulumi.Input[str] context: The path to the build context to use.
        :param pulumi.Input[str] dockerfile: The path to the Dockerfile to use.
        :param pulumi.Input[str] repository_url: Url of the repository.
        :param pulumi.Input[str] target: The target of the Dockerfile to build
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ImageArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to manage build and deployment of Docker builds. It automatically builds the Docker image and pushes it to the specified repository.

        ## Example Usage
        ### Basic Example

        ```python
        import pulumi_nuage as nuage
        
        repository = nuage.aws.Repository(
            "foo",
            name="repository",
            expire_in_days=30,
        )
        
        image = nuage.aws.Image(
            "foo",
            dockerfile="../api/Dockerfile",
            context="../",
            repository_url=repository.url,
        )
        ```

        :param str resource_name: The name of the resource.
        :param ImageArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ImageArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 architecture: Optional[pulumi.Input[str]] = None,
                 context: Optional[pulumi.Input[str]] = None,
                 dockerfile: Optional[pulumi.Input[str]] = None,
                 repository_url: Optional[pulumi.Input[str]] = None,
                 target: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.plugin_download_url is None:
            opts.plugin_download_url = _utilities.get_plugin_download_url()
        if opts.id is not None:
            raise ValueError('ComponentResource classes do not support opts.id')
        else:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ImageArgs.__new__(ImageArgs)

            __props__.__dict__["architecture"] = architecture
            __props__.__dict__["context"] = context
            if dockerfile is None and not opts.urn:
                raise TypeError("Missing required property 'dockerfile'")
            __props__.__dict__["dockerfile"] = dockerfile
            if repository_url is None and not opts.urn:
                raise TypeError("Missing required property 'repository_url'")
            __props__.__dict__["repository_url"] = repository_url
            __props__.__dict__["target"] = target
            __props__.__dict__["name"] = None
            __props__.__dict__["uri"] = None
        super(Image, __self__).__init__(
            'nuage:aws:Image',
            resource_name,
            __props__,
            opts,
            remote=True)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the docker image.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def uri(self) -> pulumi.Output[str]:
        """
        Image uri of the docker image.
        """
        return pulumi.get(self, "uri")
