<template>
  <div v-on:click="emitClickEvent" style="cursor: pointer;">
    <CWidgetDropdown
      :header=volume.driver
      :text=volume.createdAt
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
    name: 'DockerVolume',
    props: ['volume'],
    methods: {
      emitClickEvent: function() {
        this.$emit('volume-clicked', this.volume)
      }
    }
  }
</script>
