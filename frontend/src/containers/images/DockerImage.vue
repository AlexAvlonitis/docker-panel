<template>
  <div v-on:click="emitClickEvent" style="cursor: pointer;">
    <CWidgetDropdown
      :header=image.size
      :text=image.nameTag
      color="info"
    >
      <CDropdown
        class="float-right"
        color="transparent p-0"
        placement="bottom-end"
      >
        <template #toggler-content>
          <CIcon name="cil-settings"/>
        </template>
        <CDropdownItem>Action</CDropdownItem>
        <CDropdownItem>Another action</CDropdownItem>
        <CDropdownItem disabled>Disabled action</CDropdownItem>
      </CDropdown>
    </CWidgetDropdown>
  </div>
</template>

<script>
  export default {
    name: 'DockerImage',
    props: ['image'],
    methods: {
      emitClickEvent: function() {
        this.$emit('image-clicked', this.image)
      }
    }
  }
</script>
