<template>
<div class="roadmap-panel-page">
  <div
      v-for="(item,i) in roadmapItems"
      :key="i"
    >
      <h3>
        {{item.header}}
      </h3>
      <div v-html="item.description">
      </div>
  </div>
</div>
</template>

<script>
import MarkdownIt from 'markdown-it'
const md = new MarkdownIt()

export default {
  computed: {
    roadmapItems: function() {
      var roadmapItems = []
      if (this.$frontmatter && this.$frontmatter.v0) {
        roadmapItems.push({
          header: "V0: Existing or 3 months",
          description: this.$frontmatter.v0
        })
      }
      if (this.$frontmatter && this.$frontmatter.v05) {
        roadmapItems.push({
          header: "V0.5: 6 months",
          description: this.$frontmatter.v05
        })
      }
      if (this.$frontmatter && this.$frontmatter.v1) {
        roadmapItems.push({
          header: "V1: 1 Year",
          description: this.$frontmatter.v1
        })
      }
      if (this.$frontmatter && this.$frontmatter.v2) {
        roadmapItems.push({
          header: "V2: Future",
          description: this.$frontmatter.v2
        })
      }
      roadmapItems = roadmapItems.map(({header, description}) => ({header, description: `${md.render(description)}`}))
      return roadmapItems
    },
  }
}
</script>
<style scoped>
.roadmap-panel-page {
  box-sizing: border-box;
  margin-top: 1.13rem;
}</style>
