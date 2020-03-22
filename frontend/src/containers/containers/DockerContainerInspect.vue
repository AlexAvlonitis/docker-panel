<template>
  <div>
    <div v-if=loading>
      <CRow alignHorizontal=center>
        <CSpinner />
      </CRow>
    </div>
    <div v-else>
      <CRow>
        <CCol lg="12" class="mb-1">
          Filter by key: <input v-model="JSONkey"/>
        </CCol>
      </CRow>
      <pre class="p-2 code-container">{{stringifyJSON()}}</pre>
    </div>
  </div>
</template>

<script>
  import { objKeyFilter } from '@/utils/textUtils'
  import HttpClient from '@/utils/httpClient'

  export default {
    name: 'DockerContainerInspect',
    props: ['dockerContainer', 'isLoading'],
    data () {
      return {
        containerJSON: null,
        loading: this.isLoading,
        JSONkey: null
      }
    },
    methods: {
      getContainer: function(containerID) {
        HttpClient
          .get('containers', containerID)
          .then((data) => {
            this.containerJSON = data
            this.loading = false
          })
      },
      stringifyJSON: function() {
        return objKeyFilter(this.containerJSON, this.JSONkey)
      }
    },
    mounted () {
      this.getContainer(this.dockerContainer.id)
    },
    updated () {
      this.getContainer(this.dockerContainer.id)
    }
  }
</script>
