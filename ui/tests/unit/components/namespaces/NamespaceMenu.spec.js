import Vuex from 'vuex';
import { config, createLocalVue, shallowMount } from '@vue/test-utils';
import NamespaceMenu from '@/components/namespace/NamespaceMenu';

config.mocks = {
  $env: {
    isHosted: true,
  },
};

describe('NamespaceMenu', () => {
  const localVue = createLocalVue();
  localVue.use(Vuex);

  let wrapper;

  const numberNamespaces = 4;

  const namespaces = [
    {
      name: 'namespace1',
      owner: 'user1',
      member_names: ['user3', 'user4', 'user5'],
      tenant_id: 'a736a52b-5777-4f92-b0b8-e359bf484713',
    },
    {
      name: 'namespace2',
      owner: 'user1',
      member_names: ['user3', 'user4'],
      tenant_id: 'a736a52b-5777-4f92-b0b8-e359bf484714',
    },
    {
      name: 'namespace3',
      owner: 'user1',
      member_names: ['user6', 'user7', 'user8'],
      tenant_id: 'a736a52b-5777-4f92-b0b8-e359bf484715',
    },
    {
      name: 'namespace4',
      owner: 'user1',
      member_names: ['user6', 'user7'],
      tenant_id: 'a736a52b-5777-4f92-b0b8-e359bf484716',
    },
  ];

  const namespace = {
    name: 'namespace3',
    owner: 'user1',
    member_names: ['user6', 'user7', 'user8'],
    tenant_id: 'a736a52b-5777-4f92-b0b8-e359bf484715',
  };

  const store = new Vuex.Store({
    namespaced: true,
    state: {
      stats: {
        namespace,
        namespaces,
        numberNamespaces,
      },
    },
    getters: {
      'namespaces/list': (state) => state.namespaces,
      'namespaces/get': (state) => state.namespace,
    },
    actions: {
      'namespaces/fetch': () => {
      },
      'namespaces/get': () => {
      },
      'namespaces/switchNamespace': () => {
      },
      'snackbar/showSnackbarErrorLoading': () => {
      },
    },
  });

  beforeEach(() => {
    wrapper = shallowMount(NamespaceMenu, {
      store,
      localVue,
      stubs: ['fragment'],
      mocks: ['$env'],
    });
  });

  it('Is a Vue Instance', () => {
    expect(wrapper).toBeTruthy();
  });
  it('Renders the component', () => {
    expect(wrapper.html()).toMatchSnapshot();
  });
});
