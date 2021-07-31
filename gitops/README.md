# GitOps

## Basics

- GitOps is basically meaning doing Infrastructure as Code in right way.
- It all starts with Infrastructure as Code
- IaC (Infrastructure as Code) Concept:
    - when we define our Infrastructure as Code, instead of creating it manually
    - by this reproducing or replicating the Infrastructure is much more easier than doing it manually
    - Note that, Infrastructure as Code evolved as "X as Code", in the place X we can put Network, Policy, Configuration, Security
    - So it's not only Infrastructure but also others like mention above:
        - Infrastructure as Code
        - Network as Code
        - Policy as Code
        - Configuration as Code
        - Security as Code
        - so, it's basically "X as Code"
- Ex:
    - what we do by terraform, kubeform is the example of IaaS

- GitOps treats the Infrastructure as Code as the same as Application Code.
- Doing the IaC (or X as Code) using the Git repo (version control for IaC files) and the full devops pipeline for it is GitOps basically.
- How it works:
    - IaC hosted on a Git Repository
        - where it's:
            - Version controlled
            - and allows team collaboration
        - where developers open PRs on the repository (developers doesn't have the access of the Infrastructure)
        - when a PR come, it gets tested automatically by CI pipeline, it will validate the codes and test it
        - when a PR gets approve by a senior/admin then it gets merged into the main branch and then through a CD pipeline it get deployed to environment
- So by GitOps
    - we have automated the processes with more transparent
    - maintain the quality of IaC with proper tests and approval
    - team can collaborate 
- CD Pipeline:
    - Through CD pipeline changes automatically applied to the Infrastructure
    - In GitOps we have two (ways to apply these changes) types of CD pipeline, such as:
        - Push Deployement:
            - It's like our traditional application pipeline, like : jenkins, CI/CD server
            - when application is built successfully the push CD pipeline executes commands to deploy the application in the environment
        - Pull Deployment:
            - It's different than Push deployment
            - It basically do the reverse task
            - here we have an agent installed in environment ex: in k8s cluster
            - that agent actively pull the changes from our Git repository
            - agent will check regularly what is the state in the repo (desired state) and compare that with the current state of the environment where it's running
            - then it will apply the changes necessary to get to desired state
            - Tools that work with GitOps pull based model : flux CD, argo Cd
    
- So as we are using version control (git) in GitOps, so we can get back to any of our previous version by reverting. And, it will automatically sync in our environment as in GitOps we use automation (CI/CD Pipeline)
- And it has only one single sourc of truth : that's Git repo (where everything is stored), this makes easier to manage our Infrastructure
- And, It gives us security, as we don't need to give the access of our Infrastructure to all the developers, because it's the CD pipeline who deploy the changes (so only CD pipeline needs the access)
- but anyone on the team can propose changes through PRs
- we can have smaller group of people who can approve the PRs
- As a result, we have:
    - less permission to manage
    - and a more secure environment

- Summary:
    - GitOps is IaC (Infrastructure as Code) with Version Control (git) Pull requests and CI/CD pipeline.

# Resources

- [x] [What is GitOps, How GitOps works and Why it's so useful](https://www.youtube.com/watch?v=f5EpcWp0THw)