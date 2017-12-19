package org.plugins.qmstr.QuartermasterProperty

import static org.plugins.qmstr.QuartermasterProperty.DescriptorImpl.QMSTRQ_PROJECT_BLOCK_NAME

def f = namespace(lib.FormTagLib);

f.optionalBlock(name: QMSTRQ_PROJECT_BLOCK_NAME, title: _('quartermaster.project'), checked: instance != null) {
    f.entry(field:'path', title: _('qmstr_param')) {
        f.textbox()
    }
}
