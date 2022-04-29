# docker.pull
1.
Unikraft是用于构建称为专用POSIX兼容操作系统的自动化系统。 这些图像是根据特定应用程序的需求量身定制的。 Unikraft基于小型模块化库的概念,每个库均提供操作系统中常见的一部分功能(例如,内存分配,调度,文件系统支持,网络堆栈等)。 Unikraft支持多个目标平台(例如Xen,KVM和Linux用户空间),因此可以为单个应用程序构建多个映像(每个平台一个),而无需应用程序开发人员执行任何其他特定于平台的工作。 总之,Unikraft能够构建针对特定应用程序的专用OS和unikernel,而无需花费大量时间来花费大量时间来构建此类映像。

2.
虽然容器正在迅速成为行业标准，但 unikernels 仍有很多工作要做。Unikernel 试图进一步推动容器的概念，完全消除对操作系统的需求。Unikernels 通过使用库操作系统来实现这一点。库操作系统提供与常规操作系统类似的（但仅限于单用户、单地址空间）功能，但以您的应用程序使用的库的形式。因此，不是在内存中维护一个常驻内核，而是通过预先构建的二进制库来管理所有内容。但是，Unikernel 不处理资源分配，因此它们仍然需要管理程序：

3.
您可以通过将管理程序视为常规操作系统，将 unikernels 视为在其上运行的进程来了解它。不同之处在于，所有特定于应用程序的系统调用都被推送到尽可能靠近应用程序的位置，而管理程序只处理直接的硬件互操作。

可以想象，unikernel 的开销甚至比容器还要小，而且性能应该更高。此外，通过取消使用多用户、多地址空间内核，安全性得到了极大的提高。Unikernels 是一项真正令人惊叹的技术，但它们还远未准备好投入生产。首先需要解决的一些问题是：

调试。由于 unikernel 没有运行任何操作系统，因此您无法直接连接到它的外壳并进行调查。我相信会有一个更简单的方法来解决它，但还没有。
精简构建。生成 unikernel 图像很复杂，需要对该主题有深入的了解。在流程简化和标准化之前，采用将非常缓慢。
框架支持。大多数当前的应用程序框架都必须适应并生成有关 Unikernels 使用的文档。

4.
Unikraft 由三个基本组件组成：

库组件是 Unikraft 模块，每个模块都提供一些功能。正如库所期望的那样，它们是应用程序的核心构建块。库可以任意小（例如，提供概念验证调度程序的小型库）或与 libc 等标准库一样大。但是，Unikraft 中的库通常会包装预先存在的库，例如 openssl，因此现有应用程序仍然可以使用相关的现有系统，而无需重新工作。

配置。受 Linux 的 Kconfig 系统的启发，这款 Unikraft 使用这种方法让用户快速轻松地选择要包含在构建过​​程中的库，并在可用的情况下为每个库配置选项。与 Kconfig 一样，菜单跟踪依赖项并在适用时自动选择它们。

构建工具。Unikraft 的核心是一套帮助创建最终 unikernel 映像的工具。基于 make，它负责编译和链接所有相关模块，并为通过配置菜单选择的不同平台生成图像。
Library Components are Unikraft modules, each of which provides some piece of functionality. As is expected of a library, they are the core building blocks of an application. Libraries can be arbitrarily small (e.g., a small library providing a proof-of-concept scheduler) or as large as standard libraries like libc. However, libraries in Unikraft often wrap pre-existing libraries, such as openssl, and as such existing applications can still make use of relevant, existing systems without having to re-work anything.

Configuration. Inspired by Linux’s Kconfig system, this Unikraft uses this approach to quickly and easily allow users to pick and choose which libraries to include in the build process, as well as to configure options for each of them, where available. Like Kconfig, the menu keeps track of dependencies and automatically selects them where applicable.

Build Tools. The core of Unikraft is a suite of tools which aid in the creation of the final unikernel image. Based on make, it takes care of compiling and linking all the relevant modules and of producing images for the different platforms selected via the configuration menu.


5.
Nabla containers: a new approach to container isolation
Despite all of the advantages that have resulted in an industry-wide shift towards containers, containers have not been accepted as isolated sandboxes, which is crucial for container-native clouds. We introduce nabla containers, a new type of container designed for strong isolation on a host.

Nabla containers achieve isolation by adopting a strategy of attack surface reduction to the host. A visualization of this approach appears in this figure:

A containerized application can avoid making a Linux system call if it links to a library OS component that implements the system call functionality. Nabla containers use library OS (aka unikernel) techniques, specifically those from the Solo5 project, to avoid system calls and thereby reduce the attack surface. 

6.
Nabla containers only use 7 system calls; all others are blocked via a Linux seccomp policy. An overview of the internals of a nabla container appears in this figure:

For the curious, here are the allowed syscalls: read, write, exit_group, clock_gettime, ppoll, pwrite64, and pread64. They are restricted to specific file descriptors (already opened before enabling seccomp). They originate from the hypercall implementations of the ukvm unikernel monitor.1 Check out the code for more specifics.
