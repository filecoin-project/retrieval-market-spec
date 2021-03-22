<template>
  <main class="home">
    <div class="home-container theme-default-content">
      <Content class="intro" />
      <div class="grid">
        <div
          v-for="(category, i) in manualSidebar"
          :key="i"
          v-bind:class="{
            category: true,
            meta: category.title === 'Community' || category.title === 'Project'
          }"
        >
          <h2>
            <RouterLink :to="category.path" class="title">
              {{ category.title }}
            </RouterLink>
          </h2>
          <p v-for="(item, j) in category.children" :key="j">
            <RouterLink v-if="!isExternal(item.path)" :to="item.path">
              {{ item.title }}
            </RouterLink>
            <a v-else :href="item.path" target="_blank">
              {{ item.title }}
              <OutboundLink />
            </a>
          </p>
        </div>
      </div>
    </div>
  </main>
</template>

<script>
import { isExternal } from '@parent-theme/util/'

export default {
  name: 'Home',
  data: function () {
    return {
      manualSidebar: [
        {
          title: 'Exchange',
          path: '/exchange',
          children: [
            {
              title: 'Transport Protocol',
              path: '/exchange/transport',
            },
            {
              title: 'Data Transfer',
              path: '/exchange/data-tranfer',
            },
            {
              title: 'Fair Exchange',
              path: '/exchange/fair-exchange',
            },
          ]
        },
        {
          title: 'Retrieval Client',
          path: '/retrieval-client',
          children: [
            {
              title: 'API',
              path: '/retrieval-client/api'
            },
            {
              title: 'Depedencies',
              path: '/retrieval-client/dependencies'
            },
            {
              title: 'Examples',
              path: '/retrieval-client/examples'
            },
          ]
        },
        {
          title: 'Retrieval Miner',
          path: '/retrieval-miner',
          children: [
          {
            title: 'API',
            path: '/retrieval-miner/api'
          },
            {
              title: 'Depedencies',
              path: '/retrieval-miner/dependencies'
            },
            {
              title: 'Examples',
              path: '/retrieval-miner/examples'
            },
          ]
        },
        {
          title: 'Content Routing',
          path: '/content-routing',
          children: [
            {
              title: 'API',
              path: '/content-routing/api'
            },
            {
              title: 'Examples',
              path: '/content-routing/examples'
            },
          ]
        },
        {
          title: 'Payments System',
          path: '/payments',
          children: [
             {
              title: 'API',
              path: '/payments/api'
            },
            {
              title: 'Examples',
              path: '/payments/examples'
            },
          ]
        },
        {
          title: 'Content Distribution',
          path: '/content-distribution',
          children: [
            {
              title: 'API',
              path: '/content-distribution/api'
            },
            {
              title: 'Examples',
              path: '/content-distribution/examples'
            },
          ]
        },
      ]
    }
  },
  methods: {
    isExternal
  }
}
</script>

<style lang="stylus">
.home .header-anchor {
  display: none;
}
</style>

<style lang="stylus" scoped>
@media (min-width: $MQNarrow) {
  .home {$contentClass}:not(.custom) {
      background: no-repeat url("/images/main-page-background.png");
      background-position: right 3rem;
      background-size: 280px 336px;
  }
}
.home {$contentClass}:not(.custom) > h1:first-child {
    font-weight: normal;
    margin: 0 0 3rem;
}
.home {
    .intro {
	max-width: 500px;
	margin-top: 3rem;
    }
    .grid {
	margin-top: 4rem;
	display: grid;
	grid-template-columns: 1fr 1fr;
	grid-auto-flow: row dense;
	grid-auto-rows: auto;
	grid-gap: 32px;
    }
    .category {
	background: linear-gradient(-50deg, #effcf5, #e2f6f7);
	padding: 1em;
    }
    .category.meta {
	background: linear-gradient(-50deg, #dbe7f4, #e9dbf4);
    }
    .category h2 {
	font-weight: normal;
	font-size: 1.4rem;
	border-bottom: none;
	margin: 0 0 0.5rem;
    }
    .category p {
	margin: 0;
    }
    .category a {
	font-weight: normal;
    }
    .category.meta a {
	color: #5c456e;
    }

    .category a.title {
	color: black;
    }

    @media (max-width: $MQNarrow) {
	.grid {
	    grid-template-columns: 1fr;
	    /* grid-auto-rows: minmax(16rem, max-content); */
	    grid-auto-rows: auto;

	}
	.category {
	    grid-column: auto !important;
	    grid-row: auto !important;
	}

	.intro {
	  margin-top: 0;
	}
    }
}
</style>
