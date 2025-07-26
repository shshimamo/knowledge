import styles from './Modal.module.css'

type ModalProps = {
  text: string
  onYes: () => void
  onNo: () => void
}

export const Modal: React.FC<ModalProps> = (props) => {
  return (
    <div className={`${styles.modal} e2e-ui-modal`}>
      <div className={`${styles.modalContent} e2e-ui-modal-content`}>
        <h2 className="e2e-ui-modal-text">{props.text}</h2>
        <div className={styles.buttonGroup}>
          <button className={`${styles.buttonNo} e2e-ui-modal-button-no`} onClick={props.onNo}>
            No
          </button>
          <button className={`${styles.buttonYes} e2e-ui-modal-button-yes`} onClick={props.onYes}>
            Yes
          </button>
        </div>
      </div>
    </div>
  )
}
