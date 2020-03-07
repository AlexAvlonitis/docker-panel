<template>
  <div>
    <div v-if=isLoading>
      <CRow alignHorizontal=center>
        <CSpinner />
      </CRow>
    </div>
    <vue-json-pretty
      v-else
      selectableType=single
      :data=imageJSON
    />
  </div>
</template>

<script>
  import axios from 'axios';
  import VueJsonPretty from 'vue-json-pretty';

  export default {
    name: 'DockerImage',
    props: ['dockerImageID'],
    data () {
      return {
        imageJSON: null,
        isLoading: true
      }
    },
    components: {
      VueJsonPretty
    },
    mounted () {
      axios
        .get(`http://localhost:1234/images/${this.dockerImageID}`)
        .then(response => {
          this.imageJSON = response.data
          this.isLoading = false
        })
    }
  }
</script>
