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
    name: 'DockerObjectInspect',
    props: ['object', 'isLoading', 'httpEndpoint'],
    data () {
      return {
        objectJSON: null,
        loading: this.isLoading,
        JSONkey: null
      }
    },
    methods: {
      getObject: function(objID) {
        HttpClient
          .get(this.httpEndpoint, objID)
          .then((data) => {
            this.objectJSON = data
            this.loading = false
          })
      },
      stringifyJSON: function() {
        return objKeyFilter(this.objectJSON, this.JSONkey)
      }
    },
    mounted () {
      if (this.object.constructor.name === "Volume")
        return this.getObject(this.object.name)

      this.getObject(this.object.id)
    },
    updated () {
      if (this.object.constructor.name === "Volume")
        return this.getObject(this.object.name)

      this.getObject(this.object.id)
    }
  }
</script>
