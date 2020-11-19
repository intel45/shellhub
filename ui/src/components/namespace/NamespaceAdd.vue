<template>
  <fragment>
    <v-list-item-title>
      <v-dialog
        v-model="showAddNamespace"
        max-width="450"
        @click:outside="cancel"
      >
        <v-card>
          <v-card-title class="headline grey lighten-2 text-center">
            Enter Namespace
          </v-card-title>
          <ValidationObserver
            ref="obs"
            v-slot="{ passes }"
          >
            <v-card-text class="caption mb-0">
              <ValidationProvider
                v-slot="{ errors }"
                name="namespace"
                rules="required|rfc1123"
              >
                <v-text-field
                  v-model="namespace"
                  label="Namespace"
                  :error-messages="errors"
                  require
                />
              </ValidationProvider>
            </v-card-text>
            <v-card-actions>
              <v-spacer />
              <v-btn
                text
                @click="cancel"
              >
                Close
              </v-btn>
              <v-btn
                color="primary"
                text
                @click="passes(addNamespace)"
              >
                Add
              </v-btn>
            </v-card-actions>
          </ValidationObserver>
        </v-card>
      </v-dialog>
    </v-list-item-title>
  </fragment>
</template>

<script>

import {
  ValidationObserver,
  ValidationProvider,
} from 'vee-validate';

export default {
  name: 'NamespaceAdd',

  components: {
    ValidationProvider,
    ValidationObserver,
  },

  props: {
    show: {
      type: Boolean,
      required: true,
    },
  },

  data() {
    return {
      dialog: false,
      namespace: '',
    };
  },

  computed: {
    showAddNamespace: {
      get() {
        return this.show;
      },

      set(value) {
        this.$emit('show', value);
      },
    },
  },

  methods: {
    cancel() {
      this.dialog = false;
      this.$refs.obs.reset();
      this.namespace = '';
      this.$emit('update:show', false);
    },

    async addNamespace() {
      try {
        await this.$store.dispatch('namespaces/post', {
          name: this.namespace,
        });
        await this.$store.dispatch('namespaces/fetch');
        this.dialog = false;
        this.namespace = '';
        this.$emit('update:show', false);
        this.$store.dispatch('snackbar/showSnackbarSuccessAction', this.$success.namespaceCreating);
      } catch {
        this.$store.dispatch('snackbar/showSnackbarErrorAction', this.$errors.namespaceCreating);
      }
    },
  },
};

</script>
