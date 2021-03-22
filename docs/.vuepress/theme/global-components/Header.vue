<template>
  <div class="page-header">
    <h1>{{ $frontmatter.title }}</h1>
    <div v-if="confidenceLevel" class="confidence-callout">
      <h4>
        <b>Confidence Level:</b><span :class="confidenceClasses">{{ confidenceLevel }}</span>
      </h4>
    </div>
    <p>{{ $frontmatter.description }}</p>
  </div>
</template>

<script>
export default {
  computed: {
    confidenceLevel: function() {
      return this.$frontmatter && this.$frontmatter.confidenceLevel
    },

    confidenceClassName: function() {
      switch (this.confidenceLevel) {
        case 'Stable': 
          return 'stable'
        case 'Reliable':
          return 'reliable'
        case 'Draft/WIP':
          return 'draft-wip'
        case 'Brainstorm':
          return 'brainstorm'
        case 'Incorrect':
          return 'incorrect'
        case 'Missing':
          return 'missing'
        default:
          return ''
      }
    },

    confidenceClasses: function() {
      return 'confidence-level ' + this.confidenceClassName
    }
  }
}
</script>

<style lang="stylus" scoped>
.page-header {
  margin-top: -8rem !important;
  padding-top: 9.5rem !important;
}

.page-header h1 {
  margin-top: 0px;
}

.confidence-callout .confidence-level {
  margin-left: 20px
  padding: 10px
}

.confidence-callout .stable {
  background: $accentColor
  color: white
}
.confidence-callout .reliable {
  background: lighten($accentColor, 30%)
}

.confidence-callout .draft-wip {
  background: $badgeInfoColorDark
  color: white
}
.confidence-callout .brainstorm {
  background: $badgeWarningColor
}
.confidence-callout .incorrect {
  background: $badgeErrorColor
  color: white
}
.confidence-callout .missing {
  background: $codeBgColor
  color: white
}
</style>
