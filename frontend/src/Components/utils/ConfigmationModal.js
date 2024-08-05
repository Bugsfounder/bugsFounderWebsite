import React from 'react';

const ConfirmationModal = ({ isOpen, onClose, onConfirm, message }) => {
    if (!isOpen) return null;

    return (
        <div className="fixed inset-0 flex items-center justify-center bg-black bg-opacity-50">
            <div className="dark:bg-slate-700 bg-white p-6 rounded-[10px]">
                <p>{message}</p>
                <div className="flex justify-end space-x-4 mt-4">
                    <button onClick={onClose} className="px-4 py-2 bg-gray-300 text-black hover:bg-gray-200  rounded-[4px]">Cancel</button>
                    <button onClick={onConfirm} className="px-4 py-2 bg-red-500 text-white hover:bg-red-600 rounded-[4px]">Confirm</button>
                </div>
            </div>
        </div>
    );
};

export default ConfirmationModal;
