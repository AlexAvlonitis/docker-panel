<template>
  <CListGroupItem>
    {{ dockerVolume }}
  </CListGroupItem>
</template>

<script>
  export default {
    name: 'DockerVolume',
    props: ['dockerVolume']
  }
</script>
