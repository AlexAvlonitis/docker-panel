<template>
  <CListGroupItem>
    {{ dockerContainer }}
  </CListGroupItem>
</template>

<script>
  export default {
    name: 'DockerContainer',
    props: ['dockerContainer']
  }
</script>
