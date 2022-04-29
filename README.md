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


Nabla 容器：容器隔离的新方法
尽管所有这些优势都导致了整个行业向容器的转变，但容器并未被接受为孤立的沙箱，这对于容器原生云至关重要。我们介绍了nabla 容器，这是一种为主机上的强隔离而设计的新型容器。

Nabla 容器通过对主机采取减少攻击面的策略来实现隔离。此方法的可视化显示在此图中：

nabla 容器

如果容器化应用程序链接到实现系统调用功能的库操作系统组件，则它可以避免进行 Linux 系统调用。Nabla 容器使用库操作系统（又名 unikernel）技术，特别是来自Solo5 项目的技术，以避免系统调用，从而减少攻击面。Nabla 容器仅使用 7 个系统调用；所有其他人都通过 Linux seccomp 策略被阻止。下图显示了 nabla 容器的内部结构：

nabla-internals

对于好奇的人，这里是允许的系统调用：read, write, exit_group, clock_gettime, ppoll, pwrite64, 和 pread64。它们仅限于特定的文件描述符（在启用 seccomp 之前已经打开）。它们源自ukvmunikernel 监视器的超调用实现。1查看代码 了解更多详情。

nabla 容器真的更孤立吗？
nabla 容器中的隔离来自通过阻止系统调用来限制对主机内核的访问。我们通过测量容器化应用程序进行的系统调用数量以及相应地访问了多少内核函数，准确测量了 nabla 容器和标准容器对内核常见应用程序的访问量。此图总结了一些应用程序的结果：

nabla 隔离

进一步的测量和结果以及重现它们的脚本驻留在 nabla-measurements 存储库中。

存储库概述
更多信息出现在与 nabla 容器相关的每个单独的存储库中。此外， 本文将引导您完成运行第一个 nabla 容器的过程：

runnc：是 nabla 容器的 OCI 接口容器运行时。从这里开始运行 nabla 容器！

nabla-demo-apps：展示如何构建容器化为 nabla 容器的示例应用程序。有助于了解如何通过从现有的 nabla 基础 Docker 映像构建容器化您自己的应用程序。

nabla-measurements：包含 nabla 容器的隔离测量，并与标准容器和其他容器隔离方法（例如 kata 容器和 gvisor）进行比较。

如果您想更深入，请查看以下存储库：

nabla-base-build：展示如何构建 nabla 基础 Docker 镜像。有助于了解如何使用 rumprun 将应用程序或运行时移植为新的 nabla 基础。

solo5：Solo5 (http://github.com/Solo5/solo5) 的一个临时分支，其中包含“nabla-run”，这是一个基于 seccomp 的 Solo5 招标。我们正在努力在上游添加这个新的招标变更。

rumprun : Rumprun 的一个分支，使 rumprun 可以在 Solo5 界面上运行。

rumprun-packages： rumprun-packages 的一个分支，其中包含要在 Solo5 上运行的目标。

限制
主要限制是 Nabla 运行时 (runnc) 仅支持为 nabla 构建的图像（请参阅 nabla-base-build）。此处列出了其他限制 。

有关 的更多信息ukvm，请查看我们的 HotCloud '16 论文Unikernel Monitors: Extending Minimalism Outside the Box 或 Github 上的Solo5项目。 ↩

7.
Solo5 最初是由 IBM Research 的 Dan Williams 发起的一个项目，旨在移植 MirageOS 以在 Linux/KVM 管理程序上运行。从那时起，它已经发展成为一个更通用的沙盒执行环境，适合运行使用各种 unikernels（又名库操作系统）构建的应用程序，针对不同主机操作系统和虚拟机管理程序上的不同沙盒技术。

Solo5 的一些独特功能：

一个公共（“面向客户”）API，旨在轻松移植现有和未来的 unikernel-native 应用程序，
上述 API 有助于实现（“面向主机”） 绑定和招标，其设计具有隔离性、最小的攻击面 以及易于移植到不同的沙盒技术或 主机系统，
支持 unikernels 的实时和事后调试，
快速“启动”时间（相当于加载标准用户进程），适用于“功能即服务”用例。
寻找“ukvm 监视器”？自 Solo5 0.4.0 以来，我们的术语发生了变化，以更好地反映项目的预期架构和长期目标。过去被称为监视器的东西现在被称为招标。作为此更改的一部分，ukvm目标和 监视器已重命名为hvt（“硬件虚拟化招标”），以反映它们不再特定于 KVM 管理程序，并允许开发进一步的招标，例如spt。
