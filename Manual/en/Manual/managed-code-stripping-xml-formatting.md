# Link XML formatting reference

This reference provides information on the correct formatting of `link.xml` files, including valid XML elements and attributes and usage examples.

## Supported XML elements

The following XML elements are supported in link.xml files:

| **Element** | **Description** |
| --- | --- |
| `<linker></linker>` | Must be used in all `link.xml` files as the outermost element enclosing all other elements. |
| `<assembly></assembly>` | Declare code preservation annotations specific to an assembly. |
| `<type></type>` | Declare code preservation annotations specific to a type. |
| `<field></field>` | Declare code preservation annotations specific to a field. |
| `<method></method>` | Declare code preservation annotations specific to a method. |
| `<property></property>` | Declare code preservation annotations specific to a property. |
| `<event></event>` | Declare code preservation annotations specific to an event. |

For example usages of all these elements, refer to [Usage examples](#usage-examples).

## Supported XML attributes

| **Attribute** | **Description** |
| --- | --- |
| `accessors` | Apply to a `<property>` element to specify which of the property’s accessor methods to preserve. |
| `feature` | Exclude preservations for features that aren’t supported based on the settings for the current build. For more information and a usage example, refer to [Feature exclusions](#feature-exclusions). |
| `fullname` | Identify a code element by name including its parent assembly name as a prefix. Use `fullname` whenever `name` alone could be ambiguous. |
| `ignoreIfMissing` | Apply to an `<assembly>` element to declare preservations for an assembly that doesn’t exist during all Player builds. For example usage, refer to [Ignore a missing assembly](#ignore-if-missing). |
| `ignoreIfUnreferenced` | Apply to an `<assembly>` element to preserve the entities in an assembly only if at least one of its members is referenced by another assembly. For example usage, refer to [Ignore an unreferenced assembly](#ignore-if-unreferenced). |
| `name` | Identify a code element by name only, without the parent assembly as a prefix. If the assembly name has already been specified with `fullname` in a parent XML element, child XML elements can identify the code element by `name` alone. |
| `preserve` | Specify the level of code elements to preserve from stripping. |
| `signature` | Identify a code element by its signature. A method signature consists of the return type, name, and parameters. A field, property, or event signature consists of the type and name. For example usage, refer to [Usage examples](#usage-examples). |
| `windowsruntime` | Must be added to the `<assembly>` element in the `link.xml` file whenever you define preservations for a Windows Runtime Metadata (.winmd) assembly. For example usage, refer to [Windows Runtime Metadata assembly](#windows-runtime). |

For example usages of all these attributes, refer to [Usage examples](#usage-examples).

## Usage examples

The following examples demonstrate valid usage of the supported XML elements and attributes.

### Different ways to declare the root types of an assembly

The following examples illustrate the different ways that you can declare the root types of a project’s assemblies using a `link.xml` file:

```
<linker>
  <!--Preserve types and members in an assembly-->
  <assembly fullname="AssemblyName">
    <!--Preserve an entire type-->
    <type fullname="AssemblyName.TypeName" preserve="all"/>

    <!--No "preserve" attribute and no members specified means preserve all members-->
    <type fullname="AssemblyName.TypeName"/>

    <!--Preserve all fields on a type-->
    <type fullname="AssemblyName.TypeName" preserve="fields"/>

    <!--Preserve all methods on a type-->
    <type fullname="AssemblyName.TypeName" preserve="methods"/>

    <!--Preserve the type only-->
    <type fullname="AssemblyName.TypeName" preserve="nothing"/>

    <!--Preserving only specific members of a type-->
    <type fullname="AssemblyName.TypeName">
        
      <!--Fields-->
      <field signature="System.Int32 FieldName" />

      <!--Preserve a field by name rather than signature-->
      <field name="FieldName" />
      
      <!--Methods-->
      <method signature="System.Void MethodName()" />

      <!--Preserve a method with parameters-->
      <method signature="System.Void MethodName(System.Int32,System.String)" />

      <!--Preserve a method by name rather than signature-->
      <method name="MethodName" />

      <!--Properties-->

      <!--Preserve a property, its backing field (if present), 
          getter, and setter methods-->
      <property signature="System.Int32 PropertyName" />

      <property signature="System.Int32 PropertyName" accessors="all" />

      <!--Preserve a property, its backing field (if present), and getter method-->
      <property signature="System.Int32 PropertyName" accessors="get" />

      <!--Preserve a property, its backing field (if present), and setter method-->
      <property signature="System.Int32 PropertyName" accessors="set" />

      <!--Preserve a property by name rather than signature-->
      <property name="PropertyName" />

      <!--Events-->

      <!--Preserve an event, its backing field (if present), add, and remove methods-->
      <event signature="System.EventHandler EventName" />

      <!--Preserve an event by name rather than signature-->
      <event name="EventName" />

    </type>
  </assembly>
</linker>
```

### Declare entire assemblies as roots

The following example shows how you can declare entire assemblies:

```
<!--Preserve an entire assembly-->
  <assembly fullname="AssemblyName" preserve="all"/>

  <!--No "preserve" attribute and no types specified means preserve all-->
  <assembly fullname="AssemblyName"/>

  <!--Fully qualified assembly name-->
  <assembly fullname="AssemblyName, Version=0.0.0.0, Culture=neutral, PublicKeyToken=null">
    <type fullname="AssemblyName.Foo" preserve="all"/>
  </assembly>

  <!--Force an assembly to be processed for roots but don't explicitly preserve anything in particular. Useful when the assembly isn't referenced.-->
  <assembly fullname="AssemblyName" preserve="nothing"/>
```

### Preserve nested or generic types

The following example shows how to preserve either nested or generic types:

```
<!--Examples with generics-->
    <type fullname="AssemblyName.G`1">

      <!--Preserve a field with generics in the signature-->
      <field signature="System.Collections.Generic.List`1<System.Int32> FieldName" />

      <field signature="System.Collections.Generic.List`1<T> FieldName" />

      <!--Preserve a method with generics in the signature-->
      <method signature="System.Void MethodName(System.Collections.Generic.List`1<System.Int32>)" />

      <!--Preserve an event with generics in the signature-->
      <event signature="System.EventHandler`1<System.EventArgs> EventName" />

    </type>

    <!--Preserve a nested type-->
    <type fullname="AssemblyName.H/Nested" preserve="all"/>

    <!--Preserve all fields of a type if the type is used.  If the type isn't used, it will be removed-->
    <type fullname="AssemblyName.I" preserve="fields" required="0"/>

    <!--Preserve all methods of a type if the type is used. If the type isn't used, it will be removed-->
    <type fullname="AssemblyName.J" preserve="methods" required="0"/>

    <!--Preserve all types in a namespace-->
    <type fullname="AssemblyName.SomeNamespace*" preserve="all"/>

    <!--Preserve all types with a common prefix in their name-->
    <type fullname="Prefix*" preserve="all"/>
```

### Ignore a missing assembly

Use the `ignoreIfMissing` attribute on the `<assembly>` element if you need to declare preservations for an assembly that doesn’t exist during all Player builds:

```
<linker>
  <assembly fullname="Foo" ignoreIfMissing="1">
    <type name="TypeName"/>
  </assembly>
</linker>
```

### Ignore an unreferenced assembly

Use the `ignoreIfUnreferenced` attribute on the `<assembly>` element to preserve the entities in an assembly only when at least one type is referenced in another assembly:

```
<linker>
  <assembly fullname="Bar" ignoreIfUnreferenced="1">
    <type name="TypeName"/>
  </assembly>
</linker>
```

### Windows Runtime Metadata assembly

You must use the `windowsruntime` attribute on the `<assembly>` element whenever you define preservations for a Windows Runtime Metadata (.winmd) assembly:

```
<linker>
  <assembly fullname="Windows" windowsruntime="true">
    <type name="TypeName"/>
 </assembly>
</linker>
```

### Feature exclusions

The `mscorlib.xml` file embedded in `mscorlib.dll` uses this attribute, but you can use it in any `link.xml` file when appropriate.

At the **High** [stripping level](managed-code-stripping-configure.md), the Unity linker excludes preservations for features that aren’t supported based on the settings for the current build:

1. `remoting` — Excluded when targeting the **IL2CPP** **scripting backend**.
2. `sre` — Excluded when targeting the IL2CPP scripting backend.
3. `com` — Excluded when targeting platforms that don’t support COM.

For example, the following `link.xml` file preserves one method of a type on platforms that support COM, and one method on all platforms:

```
<linker>

    <assembly fullname="Foo">

        <type fullname="Type1">

            <!--Preserve FeatureOne on platforms that support COM-->

            <method signature="System.Void FeatureOne()" feature="com"/>

            <!--Preserve FeatureTwo on all platforms-->

            <method signature="System.Void FeatureTwo()"/>

        </type>

    </assembly>

</linker>
```

## Additional resources

* [Managed code stripping and the Unity linker](unity-linker.md)
* [Configure managed code stripping](managed-code-stripping-configure.md)
* [Preserving code using annotations](managed-code-stripping-preserving.md)
* [Unity linker marking rules reference](managed-code-stripping-marking-rules.md)
* [IUnityLinkerProcessor.GenerateAdditionalLinkXmlFile](../ScriptReference/Build.IUnityLinkerProcessor.GenerateAdditionalLinkXmlFile.md)