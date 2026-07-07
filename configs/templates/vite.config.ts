// @config-manager:start imports
import preact from "@preact/preset-vite";
import tailwindcss from "@tailwindcss/vite";
import { existsSync, mkdirSync, readFileSync, readdirSync, renameSync, writeFileSync } from "node:fs";
import { join, resolve } from "node:path";
import { defineConfig } from "vite";
// @config-manager:end imports

// @config-manager:start constants
const themeDir = import.meta.dirname;
const assetsDir = resolve(themeDir, "assets");
const dataDir = resolve(themeDir, "data");
const outDir = resolve(themeDir, "static/dist");
const sourcemapDir = resolve(themeDir, "sourcemaps");
// @config-manager:end constants

// @config-manager:start support_functions
/**
 * @description Custom plugin to copy the Vite-generated manifest from the
 * output directory to Hugo's data directory so templates can read it via
 * `site.Data.manifest`.
 */
function copyManifestPlugin() {
  return {
    name: "copy-manifest-to-hugo-data",
    closeBundle: function () {
      const manifestSrc = resolve(outDir, ".vite/manifest.json");

      if (existsSync(manifestSrc)) {
        mkdirSync(dataDir, { recursive: true });
        const content = readFileSync(manifestSrc, "utf-8");
        writeFileSync(resolve(dataDir, "manifest.json"), content);
      }
    },
  };
}

/**
 * @description Custom plugin to relocate source map files from static/dist/ to
 * a separate sourcemaps/ directory. This keeps deployment artifacts free of
 * debug information while preserving maps for error tracking.
 */
function relocateSourcemapsPlugin() {
  return {
    name: "relocate-sourcemaps",
    closeBundle: function () {
      mkdirSync(sourcemapDir, { recursive: true });
      moveSourcemaps(outDir);
    },
  };
}

function moveSourcemaps(dir: string) {
  const entries = readdirSync(dir, { withFileTypes: true });

  for (const entry of entries) {
    const fullPath = join(dir, entry.name);

    if (entry.isDirectory()) {
      moveSourcemaps(fullPath);
    } else if (entry.name.endsWith(".map")) {
      const dest = join(sourcemapDir, entry.name);
      renameSync(fullPath, dest);
    } else if (entry.name.endsWith(".js") || entry.name.endsWith(".css")) {
      // Remove sourceMappingURL comment since maps are relocated
      const content = readFileSync(fullPath, "utf-8");
      const updated = content.replace(/\/[/*]#\s*sourceMappingURL=.*\.map\s*\*?\/?$/gmu, "");

      if (updated !== content) {
        writeFileSync(fullPath, updated);
      }
    }
  }
}
// @config-manager:end support_functions

// @config-manager:start export
export default defineConfig({
  root: assetsDir,

  build: {
    outDir: outDir,
    emptyOutDir: true,
    manifest: true,
    sourcemap: true,

    rollupOptions: {
      input: { main: resolve(assetsDir, "css/main.css"), "main-js": resolve(assetsDir, "js/main.ts") },
      output: {
        entryFileNames: "[name].[hash:8].js",
        chunkFileNames: "[name].[hash:8].js",
        assetFileNames: function (assetInfo) {
          if (assetInfo.names?.[0]?.endsWith(".woff2") || assetInfo.names?.[0]?.endsWith(".woff")) {
            return "fonts/[name].[hash:8][extname]";
          }
          return "[name].[hash:8][extname]";
        },
        manualChunks: function (id): string | void {
          if (id.includes("node_modules/preact")) {
            return "vendor-preact";
          }
        },
      },
    },
  },

  plugins: [tailwindcss(), preact(), copyManifestPlugin(), relocateSourcemapsPlugin()],

  resolve: { alias: { "@components": resolve(assetsDir, "components") } },
});
// @config-manager:end export
