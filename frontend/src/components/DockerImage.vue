<template>
  <CListGroupItem>
    {{ dockerImage }}
  </CListGroupItem>
</template>

<script>
  export default {
    name: 'DockerImage',
    props: ['dockerImage']
  }
</script>
