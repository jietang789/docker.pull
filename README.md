https://cd.foundation/blog/2022/04/01/future-of-continuous-delivery-trends/

Ortelius简化了微服务的实现。通过提供带有部署规范的中心服务目录，应用程序团队可以轻松地跨集群使用和部署服务。Ortelius根据服务更新跟踪应用程序版本，并映射它们的服务依赖关系，从而消除混淆和猜测工作。

Kubernetes上构建容器图像的框架。


CDEvents是连续交付事件的通用规范，支持在完整的软件生产生态系统中实现互操作性

CDF项目全景图
https://landscape.cd.foundation/?grouping=organization&organization=continuous-delivery-foundation-cdf&style=borderless
CDF介绍
持续交付基金会的使命是维护并发展一个开放的持续交付生态，基金会的创始成员包括 Alauda、阿里巴巴、Anchore、Armory、Autodesk、Capital One、CircleCI、CloudBees、DeployHub、GitLab、Google、华为、JFrog、Netflix、Puppet、Red Hat、SAP 和 Snyk。


Apache 孵化项目清单链接
https://incubator.apache.org/projects/#current

Eclipse基金会项目清单链接
https://projects.eclipse.org/list-of-projects?combine=&field_project_techology_types_tid=1303&field_state_value_2=All

Tekton
https://www.infoq.cn/article/arayxto19bd6avbmxfqz
简介
Tekton 是一个功能强大且灵活的 Kubernetes 原生 CI/CD 构建框架，用于创建持续集成和交付（CI/CD）系统。 关于 Tekton ，网上可以搜到很多很多介绍文档，本文主要阐述我对 Tekton 的实现原理和背后的技术逻辑的一点理解。



Tekton 定义了 Task、TaskRun、Pipeline、PipelineRun、PipelineResource 五类核心对象，通过对 Task 和 Pipeline 的抽象，我们可以定义出任意组合的 pipeline 模板来完成各种各样的 CI/CD 任务，再通过 TaskRun、PipelineRun 和 PipelineResource 可以将这些模板套用到各个实际的项目中。



实现原理
高度抽象的结构化设计使得 Tekton 具有非常灵活的特性，那么 Tekton 是如何实现 workflow 的流转的呢？

总结
区别于传统的 CI/CD 工具（Jenkins），Tekton 是一套构建 CICD 系统的框架。 Tekton 不能使你立即获得 CI/CD 的能力。但是基于 Tekton 可以设计出各种花式的构建部署流水线。得益于 Tekton 良好的抽象，这些设计出的流水线可以作为模板在多个组织，项目间共享。Tekton 源自 Knative 的 Build-Template 项目，设计之初的一个重要目标就是使人们能够共享和重用构成 pipeline 的组件，以及 Pipeline 本身。在 Tekton 的 RoadMap 中 Tekton Catelog 就是为了实现这一目标而提出的。



区别于 Argo 这种基于 Kubernetes 的 Workflow 工具， Tekton 在工作流控制上的支持是比较弱的。一些复杂的场景比如循环，递归等都是不支持的。更不用说 Argo 在高并发和大集群调度下的性能优化。这和 Tekton 的定位有关， Tekton 定位于实现 CICD 的框架，对于 CICD 不需要过于复杂的流程控制。大部分的研发流程可以被若干个最佳实践来覆盖。而这些最佳实践应该也必须可以在不同的组织间共享，为此 Tekton 设计了 PipelineResource 的概念。 PipelineResource 是 Task 间交互的接口，也是跨平台跨组织共享重用的组件，在 PipelineResource 上还可以有很多想象空间。

Tekton 利用 Kubernetes 的 List-Watch 机制，在启动时初始化了 2 个 Controller、PipelineRunController 和 TaskRunController 。



PipelineRunController 监听 PipelineRun 对象的变化。在它的 reconcile 逻辑中，将 pipeline 中所有的 Task 构建为一张有向无环图(DAG)，通过遍历 DAG 找到当前可被调度的 Task 节点创建对应的 TaskRun 对象。



TaskRunController 监听 TaskRun 对象的变化。在它的 reconcile 逻辑中将 TaskRun 和对应 Task 转化为可执行的 Pod ，由 kubernetes 调度执行。利用 Kubernetes 的 OwnerReference 机制， PipelineRun Own TaskRun、TaskRun Own Pod、Pod 状态变更时，触发 TaskRun 的 reconcile 逻辑， TaskRun 状态变更时触发 PipelineRun 的 reconcile 逻辑。

https://tekton.dev/docs/getting-started/tasks/
