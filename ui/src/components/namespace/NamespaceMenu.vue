<template>
  <fragment>
    <v-menu
      :close-on-content-click="false"
      offset-y
    >
      <template #activator="{ on }">
        <v-chip v-on="on">
          <v-icon left>
            mdi-server
          </v-icon>
          {{ namespace.name }}
          <v-icon right>
            mdi-chevron-down
          </v-icon>
        </v-chip>
      </template>
      <v-card>
        <v-subheader>Tenant ID</v-subheader>
        <v-list-item>
          <v-list-item-content>
            <v-list-item-subtitle>
              <v-chip>
                <span data-test="tenantID-text">{{ tenant }}</span>
                <v-icon
                  v-clipboard="tenant"
                  v-clipboard:success="() => {
                    this.$store.dispatch('snackbar/showSnackbarCopy', this.$copy.tenantId);
                  }"
                  right
                >
                  mdi-content-copy
                </v-icon>
              </v-chip>
            </v-list-item-subtitle>
          </v-list-item-content>
        </v-list-item>
        <v-list>
          <v-subheader>Namespaces</v-subheader>
          <v-list-item-group>
            <v-virtual-scroll
              :height="150"
              item-height="60"
              :items="namespaces"
            >
              <template #default="{ item }">
                <v-list-item
                  :key="item.tenant_id"
                  @click="switchIn(item.tenant_id)"
                >
                  <v-list-item-icon>
                    <v-icon>mdi-login</v-icon>
                  </v-list-item-icon>
                  <v-list-item-content>
                    <v-list-item-title>
                      {{ item.name }}
                    </v-list-item-title>
                  </v-list-item-content>
                </v-list-item>
              </template>
            </v-virtual-scroll>
          </v-list-item-group>
          <v-divider />
          <v-list-item-group
            v-model="model"
            two-line
          >
            <v-list-item
              v-show="show"
              @click="dialog=!dialog"
            >
              <v-list-item-icon>
                <v-icon>mdi-plus-box</v-icon>
              </v-list-item-icon>
              <v-list-item-content>
                <v-list-item-title>
                  Create Namespace
                </v-list-item-title>
              </v-list-item-content>
            </v-list-item>
            <v-divider />
            <v-list-item
              to="/settings/namespace-manager"
            >
              <v-list-item-icon>
                <v-icon>mdi-cog</v-icon>
              </v-list-item-icon>
              <v-list-item-content>
                <v-list-item-title>
                  Settings
                </v-list-item-title>
              </v-list-item-content>
            </v-list-item>
          </v-list-item-group>
        </v-list>
      </v-card>
      <NamespaceAdd
        :show.sync="dialog"
      />
    </v-menu>
  </fragment>
</template>

<script>
import NamespaceAdd from '@/components/namespace/NamespaceAdd';

export default {
  name: 'NamespaceMenu',

  components: {
    NamespaceAdd,
  },

  data() {
    return {
      model: true,
      dialog: false,
    };
  },

  computed: {
    isOwner() {
      return this.owner === this.$store.getters['auth/id'];
    },

    owner() {
      return this.$store.getters['namespaces/get'].owner;
    },

    namespace() {
      return this.$store.getters['namespaces/get'];
    },

    namespaces() {
      return this.$store.getters['namespaces/list'];
    },

    tenant() {
      return localStorage.getItem('tenant');
    },

    show() {
      return this.$env.isHosted;
    },
  },

  watch: {
    dialog(value) {
      if (!value) {
        this.model = false;
      }
    },
  },

  created() {
    this.getNamespaces();
  },

  methods: {
    async getNamespaces() {
      try {
        // load namespaces
        await this.$store.dispatch('namespaces/get', this.tenant);
      } catch {
        this.$store.dispatch('snackbar/showSnackbarErrorLoading', this.$errors.namespaceList);
      }
    },

    async switchIn(tenant) {
      try {
        await this.$store.dispatch('namespaces/switchNamespace', {
          tenant_id: tenant,
        });
        window.location.reload();
      } catch {
        this.$store.dispatch('snackbar/showSnackbarErrorLoading', this.$errors.namespaceSwitch);
      }
    },
  },
};
</script>
