import * as React from 'react'
import { Link } from 'react-router-dom';
import { Field, reduxForm } from 'redux-form'

const HomeForm = props => {
  return (
    <form>
      <h3>電磁弁の操作</h3>
      <div>
        <button type="button" onClick={() => props.actions.setPort(1, 'on')}>開ける</button>
        <button type="button" onClick={() => props.actions.setPort(1, 'off')}>閉める</button>
      </div>

      <h3>ポート2の操作</h3>
      <div>
        <button type="button" onClick={() => props.actions.setPort(2, 'on')}>　On　</button>
        <button type="button" onClick={() => props.actions.setPort(2, 'off')}>　Off　</button>
      </div>

      <hr />
      <div>
        <Link to='/'>ログアウト</Link>
      </div>
    </form>
  )
}

export default reduxForm({
  form: 'home' // a unique identifier for this form
})(HomeForm)
