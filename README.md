# WORKER-REPORT-DOCUMENT-LINKER

Cette lambda permet de lier un document stocké dans le S3 à un rapport via un lien d'accès permanent.

# Fonctionnement

Le lien permanent du document est ajouté à la table `rptdocument` et est lié au rapport par le `proov_code`.

Ce document devient alors visible dans les pièces-jointes, par exemple depuis l'output web du rapport.

#### Cette lambda accepte un message dans ce format
```json
{
  "proov_code": "XXXXXX",
  "document": {
    "name": "name",
    "type": "pdf",
    "path": {
      "region": "region",
      "bucket": "bucket",
      "key": "path/to/file.pdf"
    }
  }
}
```

La propriété `proov_code` permet d'identifier le rapport auquel le document sera ajouté.

La propriété `document` contient les informations du document à ajouter.

  * La propriété `name` est le nom qui servira à l'affichage (e.g visible dans l'output web)
  * La propriété `type` indique le type du fichier (e.g pdf)
  * La propriété `path` contient les informations de l'endroit où est stocké le fichier dans le S3
    * La propriété `region` est la region S3 (e.g eu-west-1)
    * La propriété `bucket` est le bucket S3 (e.g production-eu)
    * La propriété `key` est le chemin d'accès vers le fichier depuis la racine du bucket S3 
