<template>
  <CCol col="12">
    <CListGroup>
      <h1>Container List</h1>
      <div v-if=isLoading >
        <CRow alignHorizontal=center>
          <CSpinner />
        </CRow>
      </div>
      <DockerContainer
        v-else
        v-for="(dockerContainer, index) in dockerContainers"
        v-bind:dockerContainer="dockerContainer"
        v-bind:index="index"
        v-bind:key="dockerContainer.Id"
      />
    </CListGroup>
  </CCol>
</template>

<script>
  import axios from 'axios';
  import DockerContainer from './DockerContainer.vue';

  export default {
    name: 'DockerContainerList',
    components: {
      DockerContainer
    },
    data () {
      return {
        dockerContainers: null,
        isLoading: false
      }
    },
    mounted () {
      axios
        .get('http://localhost:1234/containers')
        .then(response => {
          this.dockerContainers = response.data
          this.isLoading = false
        })
    }
  }
</script>
