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
https://github.com/tektoncd


argo
https://github.com/argoproj
https://argoproj.github.io/

argo workfloows使用企业:
Adobe、阿里云、贝莱德、Capital One、Data Dog、Datastax、谷歌、GitHub、IBM、Intuit、NVIDIA、SAP、New Relic和RedHat。

///
1. Heroku
Heroku 是基于Docker的容器管道。你可以在同一平台上构建，测试，验证和部署容器，而无需配置硬件或利用其他的服务提供商。
配置
Heroku应用程序使用heroku.yml清单进行配置，该清单定义了构建和部署容器所需的步骤。如下所示：
build:   
docker:     
web: Dockerfile
如果使用Git部署容器，只需运行：
$ heroku stack:set container
$ git add heroku.yml
$ git commit -m "Add heroku.yml"
$ git push heroku master
Heroku也支持管道，使你可以将容器部署到多个环境，并可以体现交付工作流程中的各个阶段的执行步骤。
好处和局限性
Heroku非常易于使用，整个管道只需要一个YAML文件。它可以受到完全管理，为测试和部署提供了多种环境，甚至在部署不当的情况下，还可以回滚更改。
但是，并非所有Heroku的功能都支持Docker部署。例如，你不能使用Heroku CI来运行应用程序的测试组件，这意味着要么在构建镜像时运行测试组件，要么使用多阶段构建。你也不能使用 pipeline promotions 将容器从一个管道阶段升级到下一个阶段。相反，你必须将容器重新部署，以达到目标阶段。
费用
Heroku 提供了一项免费计划，一个Web dyno和一个worker dyno每月有1,000小时的免费运行时间。付费计划的起价为每月每个dyno 7美元，并提供其他功能，例如更大容量的dynos和已改进的可伸缩性。有关更多信息，请参见Heroku的定价页面。
个人观点
Heroku是一个非常简单且经济高效的容器管道解决方案。它提供了完全托管的环境，使你可以完全控制CI/CD流程。有免费标准支持，值得尝试。
2. Azure DevOps
Azure DevOps是Microsoft的用于项目管理，源代码管理（SCM）和CI/CD的服务。它使你可以控制DevOps生命周期的几乎每个阶段，同时提供容器许多的高级功能，包括私有容器注册以及与Azure Kubernetes Service（AKS）的集成。Azure Pipelines提供了平台的CI/CD服务。
配置
可以使用Web界面来管理所有Azure DevOps，也可以使用YAML文件来配置Azure Pipelines。Web UI使你可以管理和跟踪部署环境以及发行版等。
好处和局限性
如果你的团队已经使用Azure，那么Azure DevOps是你现有工作流程的自然扩展。它支持托管和本地安装，还支持许多Azure部署，包括Azure App Service，Kubernetes和Azure Functions。
但是，与其他服务（包括Azure服务）集成并非易事。在Azure DevOps中，即使你使用Azure容器注册之类的服务，也需要你复制和粘贴值，使得设置变得困难。
费用
Azure Pipelines提供了一个免费服务，其中包含一个免费的并发CI/CD作业，每月1800分钟。额外的工作费用为40美元，托管（如镜像）的费用为每月每GB 2美元。若要了解更多信息，请访问Azure DevOps Services定价页面。
个人观点
Azure DevOps非常适合需要一体化DevOps解决方案，或已经使用Azure的团队。通过将其集中在一个位置中，极大地简化了开发生命周期。但是，建立起来可能很困难，并且对于只需要基本容器管道的团队来说可能过于复杂。
3. GitLab CI/CD
GitLab 从开源SCM开始，但很快发展成为完整的DevOps解决方案。与Azure DevOps一样，它提供的功能包括项目管理，私有容器注册和构建环境（包括Kubernetes）。
配置
GitLab CI/CD由GitLab Runner驱动，在自包含的环境中执行CI/CD管道中的每个步骤。可以通过gitlab-ci.yml清单完成CI/CD配置，该清单支持一些高级配置，包括逻辑条件运算和导入其他清单。
或者，你可以使用Auto DevOps无需配置即可自动化整个管道。GitLab使用Heroku buildpacks（通过Herokuish）基于源代码（在本例中为Dockerfile）自动构建应用程序。Auto DevOps可以自动运行单元测试，执行代码质量分析以及扫描镜像以查看安全性问题。
对于部署，GitLab使用dpl工具，该工具支持各种提供商，包括云平台和Kubernetes集群。
好处和局限性
GitLab提供了非常灵活的管道，你可以自行配置，也可以使用内置工具实现完全自动化。YAML配置可以更为灵活，例如创建项目依赖项以及组合来自不同项目的多个管道。由于GitLab使用现有的开源工具，例如Herokuish和dpl，因此它支持很多的项目类型，开发语言和应用部署。
尽管GitLab可以将Runners部署到现有环境中，但它本身无法维护这些环境（Google Kubernetes Engine和Amazon Elastic Kubernetes Service除外）。它还缺少像Azure Pipelines图形化的管道配置工具。
费用
GitLab提供了一个开放源代码的基本版本和一个具有附加功能的企业付费版。对于付费计划，定价基于用户数量，每月运行CI管道所花费的分钟数以及对某些功能的访问权限来划分。所有计划都包括无限容量的代码存储库，项目计划工具以及每月2,000分钟的免费管道分钟。价格从每位用户每月4美元到每位用户每月99美元不等。
个人观点
GitLab是一种功能强大的CI/CD工具，具有极为有用的功能。开源版本功能丰富，足以与许多商业选择竞争，同时还让你自托管。但是，它确实需要你维护单独的部署环境。
4. AWS Elastic Beanstalk
Elastic Beanstalk不仅仅是管道，还可以用于编排AWS资源的工具。它可以自动设置，负载均衡，扩展和监视资源，例如ECS容器，S3存储桶和EC2实例。这使你可以根据自己的特定需求在AWS内创建一个完全自定义的管道。
配置
Beanstalk配置描述了如何部署容器以及部署该容器的环境。这是在Dockerrun.aws.json文件中定义的。Beanstalk引入了独特的概念，例如：
应用程序：Beanstalk组件（如环境和版本）的逻辑集合。
应用程序版本：代码部署版本。
环境：应用程序运行所需的一组AWS资源。
好处和局限性
Beanstalk是一个非常强大的工具，不仅适用于Docker，而且适用于AWS。它提供自动伸缩，滚动更新，监控和发布管理。它还使你可以直接访问和管理资源。
但是，Beanstalk比普通管道更复杂。除非你使用单个容器环境，并且容器版本与环境紧密耦合，否则你需要在镜像仓库中预构建和托管Docker镜像。你只能通过Beanstalk CLI触发更新。因此，如果容器失败，则需要使用Beanstalk控制台手动解决它。
费用
Beanstalk本身是免费的，但是它提供的AWS组件按正常价格定价。例如，如果你使用ECS节点和ELB负载平衡器配置环境，则将向该节点和负载平衡器收费，就像你正常配置它们一样。
个人观点
有了大量可用的AWS服务，Beanstalk提供了一种管理所有服务的好方法。当用作编排工具时，它可能非常强大，但用作容器管道可能太复杂了。
5. Google Cloud Build
Cloud Build是基于Google Cloud Platform（GCP）构建的容器CI服务。它可以直接从源代码或Dockerfile构建镜像，并直接部署到GKE，Cloud Run和其他GCP服务。
配置
Cloud Build是通过cloudbuild.yaml（或JSON）文件配置的。你可以定义构建镜像的过程以及存储镜像的位置。例如，构建Docker镜像并将其推送到Google Container Registry就像运行以下命令一样简单：
name: gcr.io/cloud-builders/docker
args: ['build', '-t', 'gcr.io/$PROJECT_ID/myimage', '.']
images: ['gcr.io/$PROJECT_ID/myimage']
Cloud Build支持触发器，触发器会根据对源代码的更改自动启动构建。
好处和局限性
Cloud Build 与其他GCP服务巧妙地集成在一起，包括GKE，App Engine和Cloud Run。你可以直接控制构建计算机的大小和容量，以及镜像缓存层以加快构建速度。你还可以运行本地构建以验证或调试构建，然后再推送到Cloud Run。
由于Cloud Build是围绕GCP构建的，因此它仅支持有限数量的部署目标。可以将容器部署到其他平台，但是需要其他步骤。此外，像GitLab一样，Cloud Build也没有可视化的管道配置工具。
费用
定价基于构建机器的大小和构建时间。标准n1-standard-1实例每构建分钟的成本为0.003美元，在n1-highcpu-32实例上最高为0.064美元。在n1-standard-1实例上，你每天还可以获得120分钟的免费构建时间。
个人观点
Cloud Build相对简单，但这也是其优势之一。它快速，易学，相当便宜，并且与其他GCP服务良好集成。如果你已经有一个部署环境，或者已经使用了GCP，建议你尝试一下。
6.Jenkins X
Jenkins是最流行的CI/CD工具之一，Jenkins X通过添加全面的Kubernetes集成进一步扩展了它。Jenkins X不仅可以部署到Kubernetes，还可以为你配置和管理Kubernetes集群。
配置
Jenkins X Pipelines建立在Tekton Pipelines之上，该管道有助于在Kubernetes上运行CI/CD管道。你可以使用jenkins-x.yml文件配置管道。Jenkins X还提供了构建包，可以帮助将源代码打包到镜像中，然后将其部署到Kubernetes。
好处和局限性
Jenkins X利用两个流行的现有项目-Jenkins和Kubernetes-创建可扩展的CI/CD平台。它可以使整个CI/CD管道自动化，并支持预览环境和管道升级。因为它包括Jenkins，所以它可以访问Jenkins开发人员的整个社区。
但是，Jenkins X需要Kubernetes，对于如何配置集群有一定的规范。命令行工具可自动执行此过程的大部分操作。
费用
Jenkins X是开源的。
个人观点
对于使用Jenkins的团队来说，Jenkins X会感觉很自然。它有一些严格的限制和要求，但是对于使用Kubernetes的团队来说，拥有可以与你的基础架构集成的工具可能会有所帮助。
结论
图片

对于希望在稳定的环境中简单地部署和托管Docker容器的团队而言，Heroku很难被击败。它提供了一个快速且可配置的平台，支持广泛的集成，并拥有庞大的第三方附件市场。
Elastic Beanstalk凭借其协调AWS资源的能力紧随其后，并且支持很多复杂场景。
对于容器CI，GitLab可以说是最全面的选择，它具拥有很多功能。
Google Cloud Build利用Google Cloud Platform的速度和容量进行快速构建，并且Jenkins X受益于Jenkins项目。这些服务大多数都是开源的或提供免费试用，因此我们建议你尝试一下，看看哪种方法最适合你的工作流程。
