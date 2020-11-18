<template>
  <v-card>
    <v-dialog
      v-model="showNoNamespace"
      :retain-focus="false"
      max-width="800px"
      persistent
    >
      <v-card>
        <v-card-title
          class="headline grey lighten-2 text-center"
          primary-title
        >
          You don't have namespaces
        </v-card-title>

        <v-card-text class="mt-4 mb-0 pb-1">
          <p>
            The easiest way to configure a namespace is by using the cli
            script, as follows:
          </p>
          <p class="caption mb-0">
            Check the <a
              :href="'https://shellhub-io.github.io/'"
              target="_blank"
            >documentation</a>
            for more information and alternative install methods.
          </p>
        </v-card-text>
        <v-card-actions>
          <v-spacer />
          <v-btn
            text
            @click="close"
          >
            Close
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </v-card>
</template>

<script>

export default {
  name: 'NamespaceInstructions',

  props: {
    show: {
      type: Boolean,
      required: true,
    },
  },

  computed: {
    showNoNamespace: {
      get() {
        return this.$props.show;
      },

      set(value) {
        this.$emit('show', value);
      },
    },
  },

  methods: {
    close() {
      this.$emit('update:show', false);
      localStorage.setItem('noNamespace', true);
    },
  },
};
</script>
