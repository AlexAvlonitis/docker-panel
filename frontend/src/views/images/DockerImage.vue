<template>
  <div>
    {{dockerImage}}
  </div>
</template>

<script>
  export default {
    name: 'DockerImage',
    props: ['dockerImage']
  }
</script>
