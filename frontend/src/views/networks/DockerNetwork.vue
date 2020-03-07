<template>
  <CListGroupItem>
    {{ dockerNetwork }}
  </CListGroupItem>
</template>

<script>
  export default {
    name: 'DockerNetwork',
    props: ['dockerNetwork']
  }
</script>
