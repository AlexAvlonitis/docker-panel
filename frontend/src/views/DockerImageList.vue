<template>
  <CRow>
    <CCol lg="12" md="12" sm="12" overflow-auto>
      <h1>Images List</h1>
      <div v-if=isLoading >
        <CRow alignHorizontal=center>
          <CSpinner />
        </CRow>
      </div>
      <div v-else>
        <CDataTable
          :items="imageList()"
          :fields="fields"
          column-filter
          table-filter
          items-per-page-select
          :items-per-page="5"
          hover
          sorter
          pagination
        >
          <template #show_details="{item}">
            <td class="py-2">
              <CButton
                color="primary"
                variant="outline"
                square
                size="sm"
                @click="renderImage(item)"
              >
               Details
              </CButton>
            </td>
          </template>
        </CDataTable>
      </div>
    </CCol>
    <CCol lg="12" sm="12" md="12">
      <div v-if=selectedImage>
        <h3 class="mb-3">Details</h3>
        <DockerObjectInspect
          httpEndpoint="images"
          :object="selectedImage"
          :isLoading="true"
        />
      </div>
    </CCol>
  </CRow>
</template>

<script>
  import Image from '@/objects/image';
  import HttpClient from '@/utils/httpClient';
  import { arrayFilter } from '@/utils/textUtils'
  import DockerObjectInspect from '@/containers/DockerObjectInspect.vue';

  const fields = [
    'name',
    'id',
    'size',
    {
      key: 'show_details',
      label: '',
      _style: 'width:1%',
      sorter: false,
      filter: false
    }
  ]

  export default {
    name: 'DockerImageList',
    components: {
      DockerObjectInspect
    },
    data() {
      return {
        dockerImages: null,
        isLoading: true,
        selectedImage: null,
        imgName: null,
        fields
      }
    },
    methods: {
      imageList: function() {
        if (this.imgName == null)
          return this.dockerImages

        return arrayFilter(this.dockerImages, this.imgName)
      },
      normalizeData: function(imagesJSON) {
        return imagesJSON.map(obj => new Image(obj))
      },
      renderImage: function(img) {
        this.selectedImage = img
      }
    },
    created () {
      HttpClient
        .get('images')
        .then((data) => {
          this.dockerImages = this.normalizeData(data)
          this.isLoading = false;
        })
    }
  }
</script>
