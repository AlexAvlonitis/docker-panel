<template>
  <CListGroupItem>
    {{ dockerImage.Id }}
  </CListGroupItem>
</template>

<script>
  export default {
    name: 'DockerImage',
    props: ['dockerImage']
  }
</script>
