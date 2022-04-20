 

Goals: 
* To act as a central knowledge base for deployment events
* To abstract the need to handle Flux notifications away from interested parties
* To support pipelines that depend on the progress of Flux deployments 


Data access patterns:
* Get events for a specific revision
* Has this `revision` of this `kind` in this `namespace` been deployed?

Questions:
* What is a thing?
    * A kustomization - used without Helm releases
    * A Helm release - often a developer's main interest
    * A GitRepository - useful for informational purposes?
* If GitRepo, Kustomization and Helm release are usually related, does it make sense to fetch them by revision?
    * These resources can be named differently, how do we map them together?
* Should the API represent ordering?
    * I.e source -> kustomize -> Helm?
    * Maybe but let's just return events for now then work out how to answer questions about the model

Future ideas:
* A notification proxy for clusters that can't have direct access to Github/Slack
* Plugins for Grafana Cloud and such?
