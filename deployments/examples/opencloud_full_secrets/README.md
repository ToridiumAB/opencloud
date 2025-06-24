---
document this deployment example in: docs/opencloud/deployment/opencloud_full.md
---

# OpenCloud WOPI Deployment Example

This deployment example is documented in two locations for different audiences:

* In the [Admin Documentation](https://docs.opencloud.eu/docs/admin/intro)\
  Providing two variants using detailed configuration step by step guides:\
  [Docker Compose Setup](https://docs.opencloud.eu/docs/admin/getting-started/container/docker-compose) and [Docker Compose Local](https://docs.opencloud.eu/docs/admin/getting-started/container/docker-compose-local).\
  This variant uses certificates provided via Docker secrets.
  Configure the paths using `TRAEFIK_CERT_FILE` and `TRAEFIK_KEY_FILE` in the `.env` file.

* In the [Developer Documentation](https://docs.opencloud.eu/docs/dev/intro)\
  Providing details which are more developer focused. This description can also be used when deviating from the default.\
  Note that this examples uses custom certificates loaded from the `certs` directory and is intended for testing purposes.


