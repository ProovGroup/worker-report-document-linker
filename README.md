# WORKER-REPORT-DOCUMENT-LINKER

Cette lambda permet de lier un document stocké dans le S3 à un rapport via un lien d'accès permanent.

# Fonctionnement

Le lien permanent du document est ajouté à la table `rptdocuments` et est lié au rapport par le `proov_code`.

Ce document devient alors visible dans les pièces-jointes, par exemple depuis l'output web du rapport.

#### Cette lambda est appelé via trigger S3, le format attendu pour le document est le suivant

S3 Key: `any/path/{proov_code}/{filename}.{extension}`

Les propriétés du chemin `proov_code`, `filename` et `extension` sont utilisé en plus de la region et bucket S3 afin de lié le document au rapport

```go
type Document struct {
  Name string
  Type string
  Path Path
}

type Path struct {
  Region string
  Bucket string
  Key    string
}
```

La propriété `Name` est le nom qui servira à l'affichage (e.g visible dans l'output web)

La propriété `Type` indique le type du fichier (e.g pdf)

La propriété `Path` contient les informations de l'endroit où est stocké le fichier dans le S3
  * La propriété `Region` est la region S3 (e.g eu-west-1)
  * La propriété `Bucket` est le bucket S3 (e.g production-eu)
  * La propriété `Key` est le chemin d'accès vers le fichier depuis la racine du bucket S3 
