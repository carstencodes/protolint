# protolint dotnet port

The .NET port for protolint consists of

- a .NET global tool. This cannot be fulfilled by .NET SDK yet, 
  refer to [this issue](https://github.com/dotnet/sdk/issues/9503) for details.
- a simple MsBuild task, that will be added to .NET project and that will validate all items in the `ProtoBuf` item group.
- an MsBuild Task collections that registers protoc-gen-protolint as protoc-plugin for all `ProtoBuf` items.