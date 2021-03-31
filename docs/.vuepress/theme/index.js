// PL Docs theme

module.exports = {
  extend: '@vuepress/theme-default',
  plugins: [require('./plugins/vuepress-plugin-chunkload-redirect'), require('./plugins/vuepress-plugin-disqus')]
}
