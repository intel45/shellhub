import Vuex from 'vuex';
import { shallowMount, createLocalVue } from '@vue/test-utils';
import NamespaceInstructions from '@/components/namespace/NamespaceInstructions';

describe('NamespaceInstructions', () => {
  const localVue = createLocalVue();
  localVue.use(Vuex);

  let wrapper;
  const show = true;

  beforeEach(() => {
    wrapper = shallowMount(NamespaceInstructions, {
      localVue,
      stubs: ['fragment'],
      propsData: { show },
    });
  });

  it('Is a Vue instance', () => {
    expect(wrapper).toBeTruthy();
  });
  it('Renders the component', () => {
    expect(wrapper.html()).toMatchSnapshot();
  });
  it('Receives data in props', () => {
    expect(wrapper.vm.show).toEqual(show);
  });
});
