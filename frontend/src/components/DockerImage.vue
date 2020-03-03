<template>
  <li>
    {{ dockerImage }}
  </li>
</template>

<script>
  export default {
    name: 'DockerImage',
    props: ['dockerImage']
  }
</script>
