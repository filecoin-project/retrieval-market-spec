
module.exports = (options = {}, context) => ({
  name: 'vuepress-plugin-chunkload-redirect',
  ready() {
    const { pages } = context;
    const roadmap = { "v0": { "roles": {}, "components": {} }, "v05": {"roles": {}, "components": {}}, "v1": {"roles": {}, "components": {}}, "v2": {"roles": {}, "components": {}}}
      /**
       * 2.1. Handle frontmatter per page.
       */
      for (const { path, frontmatter } of pages) {
        if (!frontmatter || Object.keys(frontmatter).length === 0) {
          continue;
        }

        for (const key of ["v0", "v05", "v1", "v2"]) {
            const fieldValue = frontmatter[key];
            if (fieldValue !== undefined) {
              const classification = path.includes("roles") ? "roles" : "components"
              roadmap[key][classification][frontmatter.title] = fieldValue
            }
        }
      }
    context.roadmap = roadmap
  },

  clientDynamicModules() {
    return [
      {
        name: `roadmap.js`,
        content: `export default ${JSON.stringify(context.roadmap, null, 2)}`,
      },
    ]
  }
  
})
