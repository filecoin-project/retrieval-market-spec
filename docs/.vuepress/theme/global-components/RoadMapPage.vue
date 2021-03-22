<template>
<div class="roadmap-panel-page">
  <v-expansion-panels>
    <v-expansion-panel
      v-for="(item,i) in roadmapItems"
      :key="i"
    >
      <v-expansion-panel-header>
        {{item.header}}
      </v-expansion-panel-header>
      <v-expansion-panel-content v-html="item.description">
      </v-expansion-panel-content>
    </v-expansion-panel>
  </v-expansion-panels>
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
      roadmapItems = roadmapItems.map(({header, description}) => ({header, description: `<div class="v-expansion-panel-content__wrap">${md.render(description)}</div>`}))
      return roadmapItems
    },
  }
}
</script>
<style scoped>
.roadmap-panel-page {
  box-sizing: border-box;
  margin-top: 1.13rem;
}
.v-expansion-panel-header {
  font-size: 1.35rem !important;
}
</style>
<style>
.v-expansion-panel-content__wrap {
    padding: 0 24px 16px;
    flex: 1 1 auto;
    max-width: 100%;
}
</style>
<style scoped src="vuetify/dist/vuetify.min.css">