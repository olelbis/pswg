const releaseApi = "https://api.github.com/repos/olelbis/pswg/releases/latest";
const releaseUrl = "https://github.com/olelbis/pswg/releases/latest";

const labels = [
  ["darwin_arm64.tar.gz", "macOS arm64"],
  ["darwin_amd64.tar.gz", "macOS amd64"],
  ["linux_arm64.tar.gz", "Linux arm64"],
  ["linux_amd64.tar.gz", "Linux amd64"],
  ["windows_arm64.zip", "Windows arm64"],
  ["windows_amd64.zip", "Windows amd64"],
  ["linux_amd64.deb", "Debian amd64"],
  ["linux_amd64.rpm", "RPM amd64"],
];

function formatDate(value) {
  return new Intl.DateTimeFormat("en", {
    year: "numeric",
    month: "short",
    day: "numeric",
  }).format(new Date(value));
}

function assetFor(assets, suffix) {
  return assets.find((asset) => asset.name.endsWith(suffix));
}

async function hydrateLatestRelease() {
  const version = document.querySelector("#latest-version");
  const date = document.querySelector("#latest-date");
  const grid = document.querySelector("#download-grid");

  try {
    const response = await fetch(releaseApi, {
      headers: { Accept: "application/vnd.github+json" },
    });
    if (!response.ok) {
      throw new Error(`GitHub API returned ${response.status}`);
    }

    const release = await response.json();
    version.textContent = release.tag_name;
    date.textContent = `Published ${formatDate(release.published_at)}`;

    grid.replaceChildren(
      ...labels.map(([suffix, label]) => {
        const asset = assetFor(release.assets, suffix);
        const link = document.createElement("a");
        link.textContent = label;
        link.href = asset ? asset.browser_download_url : releaseUrl;
        return link;
      }),
    );
  } catch {
    version.textContent = "Latest stable release";
    date.textContent = "GitHub Releases";
  }
}

hydrateLatestRelease();
