<template>
  <div>
    <h1>Image List</h1>
    {{ info }}
  </div>
</template>

<script>
  import axios from 'axios';

  export default {
    name: 'ImageList',
    data () {
      return {
        info: null
      }
    },
    mounted () {
      axios
        .get('http://localhost:1234/images')
        .then(response => (this.info = response))
    }
  }
</script>
