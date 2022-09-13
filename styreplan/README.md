# Fork av pgbackrest_exporter 

Eneste utvidelse er eksponering av antall dager til azure token utløper.
Eksponert via prometheus endepunkt under navnet 'pgbackrest_azuretoken_dagertilutloeper'

Fork fra https://github.com/woblerr/pgbackrest_exporter/fork

# Bygging og kopiering til lederplan

* Kjør `make dist` (bygger go-programmet i en docker-kontainer)
* kopier
  `dist/artifacts/pgbackrest_exporter-...-SNAPSHOT-...-linux-x86_64.tar.gz` til
  `lederplan/lederplan-provisjonering/lederplan-postgresql-pgbackrest-repository/docker/filer`
