# KDex CLI Tools

This repository contains essential scripts and tools for the KDex ecosystem, specifically for building, packaging, and managing OCI-compliant artifacts.

## OCI Image Structure & Metadata

The `package_image` script produces OCI images that are optimized for both Kubernetes runtime usage (as `ImageVolumes`) and fast metadata retrieval via standard REST operations.

### Dual-Purpose Layering Strategy

To balance runtime performance with metadata accessibility, images are structured with overlapping layers:

1.  **Primary Filesystem Layer (`application/vnd.oci.image.layer.v1.tar+gzip`)**:
    Contains the full content of the package (e.g., Node.js modules, static assets). This includes the `importmap.json` file, ensuring that when the image is mounted as a volume in a Pod, all files are present and correctly located.

2.  **Metadata Layer (`application/json`)**:
    A standalone, uncompressed layer containing only the `importmap.json`. This layer allows external services (like UI orchestrators or CDNs) to retrieve the import map without downloading or decompressing the primary filesystem layers.

### Metadata Identifiers

The image uses specific OCI identifiers to allow for precise filtering:

-   **Artifact Type**: `application/vnd.kdex.importmap+json`
    This is set at the manifest level. It identifies the image as a KDex-compatible package containing an import map.
-   **Layer Media Type**: `application/json`
    The standalone metadata layer uses this standard media type, making it easy to identify and extract using simple JSON parsing tools like `jq`.

### Fetching Metadata via REST

Because the metadata is stored as a dedicated OCI layer, you can retrieve it using standard `curl` operations against any OCI-compliant registry (like Distribution, GitHub Packages, or Harbor) without needing specialized OCI clients.

#### Example: Fetching Import Map by Tag

```bash
# Configuration
REGISTRY="my-registry.com"
REPO="my-app/package"
TAG="latest"
BASE_URL="http://$REGISTRY/v2/$REPO"

# 1. Fetch the manifest using the tag
# 2. Extract the digest of the 'application/json' layer
# 3. Fetch that specific blob content
curl -s -L -H "Accept: application/vnd.oci.image.manifest.v1+json" \
  "$BASE_URL/manifests/$TAG" | \
  jq -r '.layers[] | select(.mediaType=="application/json") | .digest' | \
  xargs -I {} curl -s -L "$BASE_URL/blobs/{}"
```

### Compatibility Note

The image configuration (`config.json`) is automatically updated to include the `diff_id` of both layers. This ensures full compatibility with container runtimes (containerd, CRI-O), allowing the image to be used as a standard `ImageVolume` or `container` image without validation errors.
