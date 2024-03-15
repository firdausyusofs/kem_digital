import { createTheme, TextInput, Button, Tooltip, Anchor, rem, Modal, Select } from "@mantine/core";

const lfGradient = 'conic-gradient(from -45deg, #F9CDF6, #0491DE, #C8EF62, #F9CDF6)';
const greyGradient = 'linear-gradient(to right, #FFF, #F7F8FA)';


export const theme = createTheme({
  primaryColor: 'gray',
  components: {
    TextInput: TextInput.extend({
      vars: (theme, props) => {
        if (props.variant === 'lf_textinput') {
          return {
            root:{
            },
            wrapper: {
              background: '#fff',
              borderRadius: rem(30),
              border: '1px solid #E5E7EB'
            },
            input: {
              color: '#6C737F',
              borderRadius: rem(30),
            },
            label :{
              color: '#6C737F',
              fontSize: rem(14),
              fontWeight: 500
            }
          };
        }
        return { root: {} };
      },
    }),
    Select: Select.extend({
      vars: (theme, props) => {
        if (props.variant === 'lf_select') {
          return {
            root:{
            },
            wrapper: {
              background: '#fff',
              borderRadius: rem(30),
              border: '1px solid #E5E7EB'
            },
            input: {
              color: '#6C737F',
              borderRadius: rem(30),
            },
            label :{
              color: '#6C737F',
              fontSize: rem(14),
              fontWeight: 500
            }
          };
        }
        return { root: {} };
      }
    }),
    Button: Button.extend({
      //@ts-ignore
      vars: (theme, props) => {
        if (props.variant === 'lf_primarybtn') {
          return {
            root: {
              padding: rem(3),
              border: 0,
              backgroundImage: lfGradient
            },
            inner: {
              background: 'var(--mantine-color-body)',
              color: 'var(--mantine-color-text)',
              borderRadius: 'calc(var(--button-radius) - 2px)',
              paddingLeft: 'var(--mantine-spacing-md)',
              paddingRight: 'var(--mantine-spacing-md)',
              // paddingTop: rem(8),
              // paddingBottom: rem(8),
            },
            label: {
              color: '#1F2A37',
              fontWeight: 700,
              fontSize: rem(16),
              lineHeight: rem(24)
            },
          };
        } else if (props.variant === 'lf_primarybtn2') {
          return {
            root: {
              padding: rem(3),
              border: 0,
              backgroundImage: lfGradient,
              height: 'auto'
            },
            inner: {
              background: 'var(--mantine-color-body)',
              color: 'var(--mantine-color-text)',
              borderRadius: 'calc(var(--button-radius) - 2px)',
              paddingLeft: 'var(--mantine-spacing-md)',
              paddingRight: 'var(--mantine-spacing-md)',
              paddingTop: rem(6),
              paddingBottom: rem(6),
              fontWeight: 700
            },
            label: {
              fontSize: rem(16),
              color: 'rgba(31, 42, 55, 1)',
              lineHeight: rem(24)
            },
          };
        } else if (props.variant === 'lf_secondarybtn') {
          return {
            root: {
              padding: 0,
              border: '4px solid #CAEAFF',
              borderRadius: '1000px',
              height: 'fit-content',
            },
            inner: {
              background: '#F8FCFF',
              borderRadius: '1000px',
              padding: rem(8),
              height: 'fit-content',
              paddingLeft: 'var(--mantine-spacing-md)',
              paddingRight: 'var(--mantine-spacing-md)',
            },
            label: {
              color: 'var(--Grays-Gray-600, #4D5761)',

              /* Body/body 3: Bold ✍️ */
              fontFamily: 'var(--mantine-font-family)',
              fontSize: '16px',
              fontWeight: 700,
              lineHeight: '24px', /* 150% */
              display: 'flex',
              gap: rem(8),
            },
          };
        } else if (props.variant === 'lf_tertiarybtn') {
          return {
            root: {
              padding: rem(2),
              border: '1px solid #E5E7EB',
              backgroundImage: greyGradient
            },
            inner: {
              color: '#4D5761',
              borderRadius: 'calc(var(--button-radius) - 2px)',
              paddingLeft: 'var(--mantine-spacing-md)',
              paddingRight: 'var(--mantine-spacing-md)',
            },
            label: {
              color: '#4D5761',
              fontWeight: 700,
              fontSize: '16px'
            },
          };
        } else if (props.variant === 'lf_dangerbtn') {
          return {
            root: {
              padding: rem(3),
              border: 0,
              backgroundImage: 'linear-gradient(90deg, #FFC7CE 0%, #FF0000 100%)'
            },
            inner: {
              background: 'var(--mantine-color-body)',
              color: 'var(--mantine-color-text)',
              borderRadius: 'calc(var(--button-radius) - 2px)',
              paddingLeft: 'var(--mantine-spacing-md)',
              paddingRight: 'var(--mantine-spacing-md)',
              paddingTop: rem(8),
              paddingBottom: rem(8),
              fontWeight: 700
            },
            label: {
              fontSize: rem(16),
              color: 'rgba(31, 42, 55, 1)',
              lineHeight: rem(24)
            },
          };
        }
        return { root: {} };
      },
    }),
    Anchor: Anchor.extend({
      defaultProps: {
        c: '#137CBE',
      },
    }),
    Modal: Modal.extend({
      vars: (theme, props) => {
        return {
          root: {},
          content: {
            background: '#fff',
            borderRadius: rem(8),
            padding: rem(16),
            border: '1px solid D2D6DB'
          },
          title: {
            color: 'var(--Base-Black, #000)',
            fontStyle: 'normal',
          },
          overlay:{
            background: 'rgba(0, 0, 0, 0.35)',
            backdropFilter: 'blur(10px)'
          }
        };
      },
    }),
    Tooltip: Tooltip.extend({
      //@ts-ignore
      vars: (theme, props) => {
        return {
          root: {},
          tooltip: {
            borderRadius: rem(24),
            background: 'rgba(0, 0, 0, 0.60)',
            padding: `${rem(2)} ${rem(8)}`
          }
        };
      },
    })
  }
});
