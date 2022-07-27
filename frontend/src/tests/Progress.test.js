import { mount } from '@vue/test-utils'
import Progress from '../components/Progress.vue'

test('props fileName should be correct', () => {
  const wrapper = mount(Progress, {
    props: {
      fileName: 'test.csv'
    }
  })

  // Assert the rendered text of the component
  expect(wrapper.text()).toContain('test.csv')
})