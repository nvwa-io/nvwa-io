<template>
  <div class="shell-editor">
    <textarea ref="textarea"/>
  </div>
</template>

<script>
import CodeMirror from 'codemirror'
import 'codemirror/lib/codemirror.css'
// import 'codemirror/theme/rubyblue.css'
import 'codemirror/theme/blackboard.css'
import 'codemirror/mode/shell/shell'

// require('script-loader!jsonlint')
// import 'codemirror/mode/javascript/javascript'
// import 'codemirror/addon/lint/lint.css'
// import 'codemirror/addon/lint/lint'
// import 'codemirror/addon/lint/json-lint'

export default {
  name: 'ShellEditor',
  /* eslint-disable vue/require-prop-types */
  props: ['value', 'width', 'height', 'showLineNumber'],
  data() {
    return {
      shellEditor: false
    }
  },
  watch: {
    value(value) {
      const editor_value = this.shellEditor.getValue()
      if (value !== editor_value) {
        // this.shellEditor.setValue(JSON.stringify(this.value, null, 2))
        this.shellEditor.setValue(this.value)
      }
    }
  },
  mounted() {
    console.log(this.showLineNumber)
    this.shellEditor = CodeMirror.fromTextArea(this.$refs.textarea, {
      lineNumbers: !!((this.showLineNumber === 'true' || this.showLineNumber === undefined)),
      // mode: 'application/json',
      mode: 'shell',
      // gutters: ['CodeMirror-lint-markers'],
      theme: 'blackboard',
      lint: true
    })

    // this.shellEditor.setSize(!this.width ? 'auto' : this.width, !this.height ? 'auto' : this.height)

    // this.shellEditor.setValue(JSON.stringify(this.value, null, 2))
    this.shellEditor.setValue(this.value)
    this.shellEditor.on('change', cm => {
      this.$emit('changed', cm.getValue())
      this.$emit('input', cm.getValue())
    })
  },
  methods: {
    getValue() {
      return this.shellEditor.getValue()
    }
  }
}
</script>

<style scoped>
.shell-editor{
  height: 100%;
  position: relative;
}
.shell-editor >>> .CodeMirror {
  padding-top: 5px;
  height: auto;
  min-height: 180px;
  border-radius: 2px!important;
}
.shell-editor >>> .CodeMirror-scroll{
  min-height: 180px;
}
.shell-editor >>> .cm-s-rubyblue span.cm-string {
  color: #F08047;
}
.CodeMirror-gutter-wrapper {
  left: 35px!important;
}
</style>
