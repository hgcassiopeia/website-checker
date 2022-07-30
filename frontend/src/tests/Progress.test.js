import { mount } from '@vue/test-utils'
import Progress from '../components/Progress.vue'
import ResultCard from '../components/ResultCard.vue'
import DropZone from '../components/DropZone.vue'

describe('Dropzone Component', () => {
  it('renders a dropzone', () => {
    const wrapper = mount(DropZone)
    expect(wrapper.find('input[name="dropzoneFile"]').exists()).toBe(true)
  })
})

describe('Progress Component', () => {
  const wrapper = mount(Progress, {
    props: {
      fileName: 'test.csv',
      uploading: 20
    }
  })

  it('renders a progress bar', () => {
    expect(wrapper.find('#progress-bar').exists()).toBe(true)
  }),

  it('props fileName should be correct', () => {
    expect(wrapper.props().fileName).toContain('test.csv')
    expect(wrapper.props().uploading).toBe(20)
  })
})

describe('ResultCard Component', () => {
  const wrapper = mount(ResultCard, {
    props: {
      counterUp: 1,
      counterDown: 1,
      totalLinks: 2,
      processTime: "Used 800 milliseconds"
    }
  })

  it('renders a result card', () => {
    expect(wrapper.find('#result-header').exists()).toBe(true)
    expect(wrapper.find('#result-detail').exists()).toBe(true)
  }),

  it('props of ResultCard should be correct', () => {
    expect(wrapper.props().counterUp).toBe(1)
    expect(wrapper.props().counterDown).toBe(1)
    expect(wrapper.props().totalLinks).toBe(2)
    expect(wrapper.props().processTime).toContain("Used 800 milliseconds")
  })
})