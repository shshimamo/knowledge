import styles from './Modal.module.css'

type ModalProps = {
  text: string
  onYes: () => void
  onNo: () => void
}

export const Modal: React.FC<ModalProps> = (props) => {
  return (
    <div className={styles.modal}>
      <div className={styles.modalContent}>
        <h2>{props.text}</h2>
        <div className={styles.buttonGroup}>
          <button className={styles.buttonNo} onClick={props.onNo}>
            No
          </button>
          <button className={styles.buttonYes} onClick={props.onYes}>
            Yes
          </button>
        </div>
      </div>
    </div>
  )
}
