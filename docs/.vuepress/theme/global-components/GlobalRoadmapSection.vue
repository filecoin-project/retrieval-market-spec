<template>
<div class="global-roadmap-panel-page">
  <div
      v-for="(category, name) in roadmap"
      :key="name"
    >
      <h3>{{ name }}</h3>
      <div v-for="(item, itemName) in category"
      :key="itemName">
      <h4>
        {{itemName}}
      </h4>
      <div v-html="item">
      </div>
    </div>
  </div>
</div>
</template>

<script>
import MarkdownIt from 'markdown-it'
import globalRoadmap from '@dynamic/roadmap'

const md = new MarkdownIt()

const titleize = (value) => value.charAt(0).toUpperCase() + value.slice(1)

export default {
  computed: {
    roadmap: function() {
      const roadmap = globalRoadmap[this.sectionName]
      
      return Object.entries(roadmap).reduce((newRoadmap, [categoryName, category]) => 
        Object.assign(newRoadmap, { 
          [titleize(categoryName)]: Object.entries(category).reduce((newCategory, [itemName, item]) => 
            Object.assign(newCategory, {
              [itemName]: `${md.render(item)}`
            }), {})
        }), {})
    },
  },
  props: ['sectionName']
}
</script>
<style scoped>
.global-roadmap-panel-page {
  box-sizing: border-box;
  margin-top: 1.13rem;
}</style>
